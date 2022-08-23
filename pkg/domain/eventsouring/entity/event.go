package entity

import (
	"time"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

// Event represents all information about a manipulation of an aggregate.
type Event struct {
	// ID is a generated uuid for this event.
	ID vo.UUID
	// Type describes the cause of the event. (e.g. user.added)
	// it should always be in past-form
	Type vo.EventType
	// AggregateID id is the unique identifier of the aggregate
	// the client must generate it by it's own.
	AggregateID vo.UUID
	// AggregateType describes the meaning of the aggregate for this event
	// it could an object like user.
	AggregateType vo.AggregateType
	// OrgID is the organisation which owns this aggregate
	// an aggregate can only be managed by one organisation
	// use the ID of the org.
	OrgID vo.UUID
	// InstanceID is the instance where this event belongs to
	// use the ID of the instance.
	InstanceID vo.UUID
	// Version describes the definition of the aggregate at a certain point in time
	// it's used in read models to reduce the events in the correct definition.
	Version vo.Version
	// Metadata meta data.
	Metadata vo.Metadata
	// Data describe the changed fields (e.g. userName = "hodor")
	// data must always a pointer to a struct, a struct or a byte array containing json bytes
	Data []byte
	// Sequence is the sequence of the event.
	Sequence uint64
	// PreviousAggregateSequence is the sequence of the previous sequence of the aggregate, (e.g. org.250989)
	// if it's 0 then it's the first event of this aggregate.
	PreviousAggregateSequence uint64
	// PreviousAggregateTypeSequence is the sequence of the previous sequence of the aggregate root, (e.g. org)
	// the first event of the aggregate has previous aggregate root sequence 0.
	PreviousAggregateTypeSequence uint64
	// Service should be a unique identifier for the service which created the event,
	// it's meant for maintainability.
	Service string
	// Creator should be a unique identifier for the user which created the event,
	// it's meant for maintainability,
	// It's recommend to use the aggregate id of the user.
	Creator vo.UUID
	// CreateTime is the time the event is created
	// it's used for human readability.
	// Don't use it for event ordering,
	// time drifts in different services could cause integrity problems.
	CreateTime time.Time
}

func (e Event) GetID() vo.UUID {
	return e.ID
}

func (e Event) GetType() vo.EventType {
	return e.Type
}

func (e Event) GetAggregate() Aggregate {
	return Aggregate{
		ID:         e.AggregateID,
		Type:       e.AggregateType,
		OrgID:      e.OrgID,
		InstanceID: e.InstanceID,
		Version:    e.Version,
	}
}

func (e Event) GetMetadata() vo.Metadata {
	return e.Metadata
}

func (e Event) GetData() []byte {
	return e.Data
}

func (e Event) GetSequence() uint64 {
	return e.Sequence
}

func (e Event) GetPreviousAggregateSequence() uint64 {
	return e.PreviousAggregateSequence
}

func (e Event) GetPreviousAggregateTypeSequence() uint64 {
	return e.PreviousAggregateTypeSequence
}

func (e Event) GetService() string {
	return e.Service
}

func (e Event) GetCreator() vo.UUID {
	return e.Creator
}

func (e Event) GetCreateTime() time.Time {
	return e.CreateTime
}

func (e Event) String() string {
	return string(e.AggregateType) + "/" + string(e.Type)
}

// Validate validate event.
func (e Event) Validate() error {
	if err := e.GetType().Validate(); err != nil {
		return err
	}
	if err := e.GetAggregate().Validate(); err != nil {
		return err
	}
	return nil
}
