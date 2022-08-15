package stmt

import (
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/aggregate/projection"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

// reduce implements handler.Reduce function
func (h *StatementHandler) reduce(event event.IEvent) (*vo.Statement, error) {
	reduce, ok := h.reduces[event.GetType()]
	if !ok {
		return projection.NewNoOpStatement(event), nil
	}

	return reduce(event)
}
