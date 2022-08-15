package projection

import (
	"github.com/galaxyobe/go-ddd/pkg/domain/database/interfaces"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

// EventReducer represents the required data to work with events
type EventReducer struct {
	Event  vo.EventType
	Reduce Reduce
}

// AggregateReducer represents aggregates EventReducer
type AggregateReducer struct {
	Aggregate     vo.AggregateType
	EventReducers []EventReducer
}

type Check struct {
	Executes []func(ex interfaces.IExecContext, projectionName string) (bool, error)
}

func (c *Check) IsNoop() bool {
	return len(c.Executes) == 0
}
