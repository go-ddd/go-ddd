package entity

import (
	. "github.com/galaxyobe/go-ddd/pkg/domain/eventstore/types"

	"github.com/galaxyobe/go-ddd/pkg/constraints"
)

// Aggregate is the basic implementation of Aggregater
type Aggregate[GUID constraints.GUID] struct {
	// ID is the unique identifier of this aggregate
	ID GUID `json:"-"`
	// Type is the name of the aggregate.
	Type AggregateType `json:"-"`
	// Owner is the aggregates belongs to
	Owner GUID `json:"-"`
	// Version is the semver this aggregate represents
	Version Version `json:"-"`
}
