package aggregate

import (
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

// EventReducer represents the required data to work with events
type EventReducer struct {
	Event  vo.EventType
	Reduce Reduce
}

// AggregateReducer represents the required data to work with aggregates
type AggregateReducer struct {
	Aggregate     vo.AggregateType
	EventReducers []EventReducer
}
