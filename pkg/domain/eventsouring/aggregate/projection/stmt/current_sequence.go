package stmt

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

const (
	currentSequenceStmtFormat        = `SELECT current_sequence, aggregate_type, instance_id FROM %s WHERE projection_name = $1 AND instance_id = ANY ($2) FOR UPDATE`
	updateCurrentSequencesStmtFormat = `UPSERT INTO %s (projection_name, aggregate_type, current_sequence, instance_id, timestamp) VALUES `
)

type currentSequences map[vo.AggregateType][]*instanceSequence

type instanceSequence struct {
	instanceID string
	sequence   uint64
}

func (h *StatementHandler) currentSequences(ctx context.Context, query func(context.Context, string, ...interface{}) (*sql.Rows, error), instanceIDs []string) (currentSequences, error) {
	rows, err := query(ctx, h.currentSequenceStmt, h.ProjectionName, vo.StringArray(instanceIDs))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sequences := make(currentSequences, len(h.aggregates))
	for rows.Next() {
		var (
			aggregateType vo.AggregateType
			sequence      uint64
			instanceID    string
		)

		err = rows.Scan(&sequence, &aggregateType, &instanceID)
		if err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}

		sequences[aggregateType] = append(sequences[aggregateType], &instanceSequence{
			sequence:   sequence,
			instanceID: instanceID,
		})
	}

	if err = rows.Close(); err != nil {
		return nil, errors.Wrap(err, "close rows failed")
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "errors in scanning rows")
	}

	return sequences, nil
}

func (h *StatementHandler) updateCurrentSequences(tx *sql.Tx, sequences currentSequences) error {
	valueQueries := make([]string, 0, len(sequences))
	valueCounter := 0
	values := make([]interface{}, 0, len(sequences)*3)
	for aggregate, instanceSequence := range sequences {
		for _, sequence := range instanceSequence {
			valueQueries = append(valueQueries, "($"+strconv.Itoa(valueCounter+1)+", $"+strconv.Itoa(valueCounter+2)+", $"+strconv.Itoa(valueCounter+3)+", $"+strconv.Itoa(valueCounter+4)+", NOW())")
			valueCounter += 4
			values = append(values, h.ProjectionName, aggregate, sequence.sequence, sequence.instanceID)
		}
	}

	res, err := tx.Exec(h.updateSequencesBaseStmt+strings.Join(valueQueries, ", "), values...)
	if err != nil {
		return errors.Wrap(err, "unable to exec update sequence")
	}
	if rows, _ := res.RowsAffected(); rows < 1 {
		return errSeqNotUpdated
	}
	return nil
}
