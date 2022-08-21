package stmt

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/aggregate/projection"
)

// Init implements handler.Init
func (h *StatementHandler) Init(ctx context.Context, checks ...*projection.Check) error {
	logger := h.Logger.WithContext(ctx)
	for _, check := range checks {
		if check == nil || check.IsNoop() {
			return nil
		}
		tx, err := h.client.BeginTx(ctx, nil)
		if err != nil {
			return errors.New("begin failed")
		}
		for i, execute := range check.Executes {
			logger.Debug("executing check", zap.Int("execute", i))
			next, err := execute(h.client, h.ProjectionName)
			if err != nil {
				tx.Rollback()
				return err
			}
			if !next {
				logger.Debug("skipping next check", zap.Int("execute", i))
				break
			}
		}
		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}
