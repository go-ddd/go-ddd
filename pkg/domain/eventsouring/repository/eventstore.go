package repository

import (
	"context"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/entity"

	"github.com/Masterminds/squirrel"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

// IEventStore is an interface for an event sourcing event store.
type IEventStore interface {
	// Health checks if the connection to the storage is available
	Health(ctx context.Context) error
	// Push appends all events in the event stream to the store.
	Push(ctx context.Context, events []*entity.Event, uniqueConstraints ...*vo.UniqueConstraint) error
	// Filter returns all events matching the given search query.
	Filter(ctx context.Context, query squirrel.SelectBuilder) (events []*entity.Event, err error)
	// LatestSequence returns the latest sequence.
	LatestSequence(ctx context.Context, query squirrel.SelectBuilder) (uint64, error)
	// InstanceIDs returns the instance ids found by the search query
	InstanceIDs(ctx context.Context, query squirrel.SelectBuilder) ([]string, error)
	// CreateInstance creates a new sequence for the given instance
	CreateInstance(ctx context.Context, instanceID string) error
	// Close closes the EventStore.
	Close() error
}
