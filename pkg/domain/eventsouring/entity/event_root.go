package entity

import (
	"context"
	"time"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/contexts"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

type EventRoot struct {
	EventType                     vo.EventType `json:"-"`
	Aggregate                     Aggregate    `json:"-"`
	Metadata                      vo.Metadata  `json:"-"`
	Data                          vo.EventData `json:"-"`
	Sequence                      uint64       `json:"-"`
	PreviousAggregateSequence     uint64       `json:"-"`
	PreviousAggregateTypeSequence uint64       `json:"-"`
	Service                       string       `json:"-"`
	Creator                       vo.GUID      `json:"-"`
	CreateTime                    time.Time    `json:"-"`
}

func (e EventRoot) GetType() vo.EventType {
	return e.EventType
}

func (e EventRoot) GetAggregate() Aggregate {
	return e.Aggregate
}

func (e EventRoot) GetMetadata() vo.Metadata {
	return e.Metadata
}

func (e EventRoot) GetData() vo.EventData {
	return e.Data
}

func (e EventRoot) GetSequence() uint64 {
	return e.Sequence
}

func (e EventRoot) GetPreviousAggregateSequence() uint64 {
	return e.PreviousAggregateSequence
}

func (e EventRoot) GetPreviousAggregateTypeSequence() uint64 {
	return e.PreviousAggregateTypeSequence
}

func (e EventRoot) GetService() string {
	return e.Service
}

func (e EventRoot) GetCreator() vo.GUID {
	return e.Creator
}

func (e EventRoot) GetCreateTime() time.Time {
	return e.CreateTime
}

func (e EventRoot) String() string {
	return string(e.Aggregate.Type) + "/" + string(e.EventType)
}

// Validate validate event.
func (e EventRoot) Validate() error {
	if err := e.GetType().Validate(); err != nil {
		return err
	}
	if err := e.GetAggregate().Validate(); err != nil {
		return err
	}
	return nil
}

// NewEventRootFromEvent maps a stored event to a EventRoot
func NewEventRootFromEvent(event *Event) *EventRoot {
	return &EventRoot{
		EventType: event.Type,
		Aggregate: Aggregate{
			ID:         event.AggregateID,
			Type:       event.AggregateType,
			OrgID:      event.OrgID,
			InstanceID: event.InstanceID,
			Version:    event.Version,
		},
		Metadata:                      event.Metadata,
		Data:                          vo.NewBytesEventData(event.Data),
		Sequence:                      event.Sequence,
		PreviousAggregateSequence:     event.PreviousAggregateSequence,
		PreviousAggregateTypeSequence: event.PreviousAggregateTypeSequence,
		Service:                       event.Service,
		Creator:                       event.Creator,
		CreateTime:                    event.CreateTime,
	}
}

// NewEventRootForPush is the constructor for event's which will be pushed into the eventstore
// the org id of the aggregate is only used if it's the first event of this aggregate type
// afterwards the resource owner of the first previous events is taken
func NewEventRootForPush(ctx context.Context, aggregate *Aggregate, typ vo.EventType) *EventRoot {
	return &EventRoot{
		EventType: typ,
		Aggregate: *aggregate,
		Service:   contexts.GetService(ctx),
		Creator:   contexts.GetCreator(ctx),
	}
}
