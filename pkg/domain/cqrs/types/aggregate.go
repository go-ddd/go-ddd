package types

// AggregateType is the type of aggregate.
type AggregateType string

// String returns the string representation of an event type.
func (at AggregateType) String() string {
	return string(at)
}
