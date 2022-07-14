package types

// EventType is the type of event, used as its unique identifier.
type EventType string

// String returns the string representation of an event type.
func (et EventType) String() string {
	return string(et)
}

// EventData is any additional data for an event.
type EventData interface{}
