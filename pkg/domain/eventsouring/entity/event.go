package entity

import (
	"time"

	"github.com/galaxyobe/go-ddd/pkg/constraints"
)

// Event is a domain event describing a change that has happened to an aggregate.
type Event[GUID constraints.GUID] interface {
	// Type returns the type of the event.
	Type() EventType
	// Data The data attached to the event.
	Data() EventData
	// Timestamp of when the event was created.
	Timestamp() time.Time
	// Service is the service which pushed the event
	Service() GUID
	// User is the user who pushed the event
	User() GUID
	// Metadata is app-specific metadata such as request ID, originating user etc.
	Metadata() map[string]any
	// Aggregate returns the aggregate.
	Aggregate() Aggregate[GUID]
	// String A string representation of the event.
	String() string
}
