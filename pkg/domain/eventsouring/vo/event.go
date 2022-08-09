package vo

import (
	"errors"
)

// AggregateType is the type of aggregate.
type AggregateType string

// String returns the string representation of an event type.
func (at AggregateType) String() string {
	return string(at)
}

// EventType is the type of event, used as its unique identifier.
type EventType string

// String returns the string representation of an event type.
func (et EventType) String() string {
	return string(et)
}

func (et EventType) Validate() error {
	if et == "" {
		return errors.New("event type must not be empty")
	}
	return nil
}

// EventData is any additional data for an event.
type EventData interface {
	MarshalData() ([]byte, error)
	UnmarshalData([]byte) error
}

type Bytes struct {
	data []byte
}

func NewBytesEventData(data []byte) EventData {
	return &Bytes{
		data: data,
	}
}

func (b Bytes) MarshalData() ([]byte, error) {
	return b.data, nil
}

func (b Bytes) UnmarshalData(bytes []byte) error {
	b.data = bytes
	return nil
}
