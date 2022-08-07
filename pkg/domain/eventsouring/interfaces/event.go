package interfaces

import (
	"time"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/entity"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

// ICommand is a domain command describing a change that has happened to an aggregate.
type ICommand interface {
	// GetType returns the type of the event.
	GetType() vo.EventType
	// GetAggregate returns the aggregate.
	GetAggregate() entity.Aggregate
	// GetMetadata return app-specific metadata such as request ID, originating user etc.
	GetMetadata() vo.Metadata
	// GetData return the data attached to the event.
	GetData() vo.EventData
	// GetService return the service which pushed the event.
	GetService() string
	// GetCreator return the user who pushed the event.
	GetCreator() vo.GUID
	// GetUniqueConstraints return command unique constraints.
	GetUniqueConstraints() []*vo.UniqueConstraint
	// Validate validate event.
	Validate() error
	// String A string representation of the event.
	String() string
}

// IEvent is a domain event describing a change that has happened to an aggregate.
type IEvent interface {
	// GetID return a generated uuid for this event
	GetID() vo.GUID
	// GetType returns the type of the event.
	GetType() vo.EventType
	// GetAggregate returns the aggregate.
	GetAggregate() entity.Aggregate
	// GetMetadata return app-specific metadata such as request ID, originating user etc.
	GetMetadata() vo.Metadata
	// GetData return the data attached to the event.
	GetData() []byte
	// GetSequence return the Event sequence.
	GetSequence() uint64
	// GetPreviousAggregateSequence returns the previous sequence of the aggregate root (e.g. for org.42508134)
	GetPreviousAggregateSequence() uint64
	// GetPreviousAggregateTypeSequence returns the previous sequence of the aggregate type (e.g. for org)
	GetPreviousAggregateTypeSequence() uint64
	// GetService return the service which pushed the event.
	GetService() string
	// GetCreator return the user who pushed the event.
	GetCreator() vo.GUID
	// GetCreateTime return when the event was created.
	GetCreateTime() time.Time
	// String A string representation of the event.
	String() string
}
