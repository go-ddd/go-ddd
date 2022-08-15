package stmt

import (
	"database/sql"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

const (
	setFailureCountStmtFormat = "UPSERT INTO %s" +
		" (projection_name, failed_sequence, failure_count, error, instance_id)" +
		" VALUES ($1, $2, $3, $4, $5)"
	failureCountStmtFormat = "WITH failures AS (SELECT failure_count FROM %s WHERE projection_name = $1 AND failed_sequence = $2 AND instance_id = $3)" +
		" SELECT IF(" +
		"EXISTS(SELECT failure_count FROM failures)," +
		" (SELECT failure_count FROM failures)," +
		" 0" +
		") AS failure_count"
)

func (h *StatementHandler) handleFailedStmt(tx *sql.Tx, stmt *vo.Statement, execErr error) (shouldContinue bool) {
	failureCount, err := h.failureCount(tx, stmt.Sequence, stmt.InstanceID)
	if err != nil {
		h.Logger.Warn("unable to get failure count",
			zap.Uint64("sequence", stmt.Sequence),
			zap.Error(err),
		)
		return false
	}
	failureCount += 1
	err = h.setFailureCount(tx, stmt.Sequence, failureCount, execErr, stmt.InstanceID)
	h.Logger.OnError(err, zap.Uint64("sequence", stmt.Sequence)).Warn("unable to update failure count\"")
	return failureCount >= h.maxFailureCount
}

func (h *StatementHandler) failureCount(tx *sql.Tx, seq uint64, instanceID string) (count uint, err error) {
	row := tx.QueryRow(h.failureCountStmt, h.ProjectionName, seq, instanceID)
	if err = row.Err(); err != nil {
		return 0, errors.Wrap(err, "unable to update failure count")
	}
	if err = row.Scan(&count); err != nil {
		return 0, errors.Wrap(err, "unable to scan count")
	}
	return count, nil
}

func (h *StatementHandler) setFailureCount(tx *sql.Tx, seq uint64, count uint, err error, instanceID string) error {
	_, dbErr := tx.Exec(h.setFailureCountStmt, h.ProjectionName, seq, count, err.Error(), instanceID)
	if dbErr != nil {
		return errors.Wrap(dbErr, "set failure count failed")
	}
	return nil
}
