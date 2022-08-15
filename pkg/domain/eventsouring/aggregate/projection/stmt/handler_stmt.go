package stmt

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/aggregate/projection"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

var (
	errSeqNotUpdated = errors.New("current sequence not updated")
)

type StatementHandlerOptions struct {
	projection.HandlerOptions

	Client            *sql.DB
	SequenceTable     string
	LockTable         string
	FailedEventsTable string
	MaxFailureCount   uint
	BulkLimit         uint64

	Reducers  []projection.AggregateReducer
	InitCheck *projection.Check
}

type StatementHandler struct {
	*projection.Handler
	Locker

	client                  *sql.DB
	sequenceTable           string
	currentSequenceStmt     string
	updateSequencesBaseStmt string
	maxFailureCount         uint
	failureCountStmt        string
	setFailureCountStmt     string

	aggregates []vo.AggregateType
	reduces    map[vo.EventType]projection.Reduce

	bulkLimit uint64
}

func NewStatementHandler(
	ctx context.Context,
	options StatementHandlerOptions,
) StatementHandler {
	aggregateTypes := make([]vo.AggregateType, 0, len(options.Reducers))
	reduces := make(map[vo.EventType]projection.Reduce, len(options.Reducers))
	for _, aggReducer := range options.Reducers {
		aggregateTypes = append(aggregateTypes, aggReducer.Aggregate)
		for _, eventReducer := range aggReducer.EventReducers {
			reduces[eventReducer.Event] = eventReducer.Reduce
		}
	}

	h := StatementHandler{
		client:                  options.Client,
		sequenceTable:           options.SequenceTable,
		maxFailureCount:         options.MaxFailureCount,
		currentSequenceStmt:     fmt.Sprintf(currentSequenceStmtFormat, options.SequenceTable),
		updateSequencesBaseStmt: fmt.Sprintf(updateCurrentSequencesStmtFormat, options.SequenceTable),
		failureCountStmt:        fmt.Sprintf(failureCountStmtFormat, options.FailedEventsTable),
		setFailureCountStmt:     fmt.Sprintf(setFailureCountStmtFormat, options.FailedEventsTable),
		aggregates:              aggregateTypes,
		reduces:                 reduces,
		bulkLimit:               options.BulkLimit,
		Locker:                  NewLocker(options.Client, options.LockTable, options.HandlerOptions.ProjectionName),
	}
	h.Handler = projection.NewProjectionHandler(ctx, options.HandlerOptions, h.reduce, h.Update, h.SearchQuery, h.Lock, h.Unlock)

	err := h.Init(ctx, options.InitCheck)
	h.Logger.OnError(err).Fatal("unable to initialize projections")

	h.Subscribe(h.aggregates...)

	return h
}

func (h *StatementHandler) SearchQuery(ctx context.Context, instanceIDs []string) (squirrel.SelectBuilder, uint64, error) {
	sequences, err := h.currentSequences(ctx, h.client.QueryContext, instanceIDs)
	if err != nil {
		return squirrel.SelectBuilder{}, 0, err
	}

	// queryBuilder := vo.NewSearchQueryBuilder(vo.ColumnsEvent).Limit(h.bulkLimit)

	for _, aggregateType := range h.aggregates {
		for _, instanceID := range instanceIDs {
			var seq uint64
			for _, sequence := range sequences[aggregateType] {
				if sequence.instanceID == instanceID {
					seq = sequence.sequence
					break
				}
			}
			_ = seq
			// queryBuilder.
			// 	AddQuery().
			// 	AggregateTypes(aggregateType).
			// 	SequenceGreater(seq).
			// 	InstanceID(instanceID)
		}
	}

	return squirrel.SelectBuilder{}, h.bulkLimit, nil
}

// Update implements projection.Update
func (h *StatementHandler) Update(ctx context.Context, stmts []*vo.Statement, reduce projection.Reduce) (index int, err error) {
	if len(stmts) == 0 {
		return -1, nil
	}
	instanceIDs := make([]string, 0, len(stmts))
	for _, stmt := range stmts {
		instanceIDs = appendToInstanceIDs(instanceIDs, stmt.InstanceID)
	}
	tx, err := h.client.BeginTx(ctx, nil)
	if err != nil {
		return -1, errors.Wrap(err, "begin failed")
	}

	sequences, err := h.currentSequences(ctx, tx.QueryContext, instanceIDs)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	// checks for events between create statement and current sequence
	// because there could be events between current sequence and a creation event
	// and we cannot check via stmt.PreviousSequence
	if stmts[0].PreviousSequence == 0 {
		previousStmts, err := h.fetchPreviousStmts(ctx, tx, stmts[0].Sequence, stmts[0].InstanceID, sequences, reduce)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		stmts = append(previousStmts, stmts...)
	}

	lastSuccessfulIdx := h.executeStmts(ctx, tx, &stmts, sequences)

	if lastSuccessfulIdx >= 0 {
		err = h.updateCurrentSequences(tx, sequences)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	if lastSuccessfulIdx < len(stmts)-1 {
		return lastSuccessfulIdx, projection.ErrSomeStmtsFailed
	}

	return lastSuccessfulIdx, nil
}

func (h *StatementHandler) fetchPreviousStmts(ctx context.Context, tx *sql.Tx, stmtSeq uint64, instanceID string, sequences currentSequences, reduce projection.Reduce) (previousStmts []*vo.Statement, err error) {
	// query := vo.NewSearchQueryBuilder(vo.ColumnsEvent).SetTx(tx)
	queriesAdded := false
	for _, aggregateType := range h.aggregates {
		for _, sequence := range sequences[aggregateType] {
			if stmtSeq <= sequence.sequence && instanceID == sequence.instanceID {
				continue
			}

			// query.
			// 	AddQuery().
			// 	AggregateTypes(aggregateType).
			// 	SequenceGreater(sequence.sequence).
			// 	SequenceLess(stmtSeq).
			// 	InstanceID(sequence.instanceID)

			queriesAdded = true
		}
	}

	if !queriesAdded {
		return nil, nil
	}

	events, err := h.Eventstore.Filter(ctx, squirrel.SelectBuilder{})
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		stmt, err := reduce(event)
		if err != nil {
			return nil, err
		}
		previousStmts = append(previousStmts, stmt)
	}
	return previousStmts, nil
}

func (h *StatementHandler) executeStmts(
	ctx context.Context,
	tx *sql.Tx,
	stmts *[]*vo.Statement,
	sequences currentSequences,
) int {

	lastSuccessfulIdx := -1
stmts:
	for i := 0; i < len(*stmts); i++ {
		stmt := (*stmts)[i]
		for _, sequence := range sequences[stmt.AggregateType] {
			if stmt.Sequence <= sequence.sequence && stmt.InstanceID == sequence.instanceID {
				// logging.WithFields("statement", stmt, "currentSequence", sequence).Debug("statement dropped")
				if i < len(*stmts)-1 {
					copy((*stmts)[i:], (*stmts)[i+1:])
				}
				*stmts = (*stmts)[:len(*stmts)-1]
				i--
				continue stmts
			}
			if stmt.PreviousSequence > 0 && stmt.PreviousSequence != sequence.sequence && stmt.InstanceID == sequence.instanceID {
				// logging.WithFields("projection", h.ProjectionName, "aggregateType", stmt.AggregateType, "sequence", stmt.Sequence, "prevSeq", stmt.PreviousSequence, "currentSeq", sequence.sequence).Warn("sequences do not match")
				break stmts
			}
		}
		err := h.executeStmt(ctx, tx, stmt)
		if err == nil {
			updateSequences(sequences, stmt)
			lastSuccessfulIdx = i
			continue
		}

		shouldContinue := h.handleFailedStmt(tx, stmt, err)
		if !shouldContinue {
			break
		}

		updateSequences(sequences, stmt)
		lastSuccessfulIdx = i
		continue
	}
	return lastSuccessfulIdx
}

// executeStmt handles sql statements
// an error is returned if the statement could not be inserted properly
func (h *StatementHandler) executeStmt(ctx context.Context, tx *sql.Tx, stmt *vo.Statement) error {
	if stmt.IsNoop() {
		return nil
	}
	_, err := tx.Exec("SAVEPOINT push_stmt")
	if err != nil {
		return errors.Wrap(err, "unable to create savepoint")
	}
	err = stmt.Execute(ctx, tx, h.ProjectionName)
	if err != nil {
		_, rollbackErr := tx.Exec("ROLLBACK TO SAVEPOINT push_stmt")
		if rollbackErr != nil {
			return errors.Wrap(rollbackErr, "rollback to savepoint failed")
		}
		return errors.Wrap(err, "unable execute stmt")
	}
	_, err = tx.Exec("RELEASE push_stmt")
	if err != nil {
		return errors.Wrap(err, "unable to release savepoint")
	}
	return nil
}

func updateSequences(sequences currentSequences, stmt *vo.Statement) {
	for _, sequence := range sequences[stmt.AggregateType] {
		if sequence.instanceID == stmt.InstanceID {
			sequence.sequence = stmt.Sequence
			return
		}
	}
	sequences[stmt.AggregateType] = append(sequences[stmt.AggregateType], &instanceSequence{
		instanceID: stmt.InstanceID,
		sequence:   stmt.Sequence,
	})
}

func appendToInstanceIDs(instances []string, id string) []string {
	for _, instance := range instances {
		if instance == id {
			return instances
		}
	}
	return append(instances, id)
}
