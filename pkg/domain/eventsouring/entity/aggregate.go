package entity

import (
	"errors"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

// Aggregate is the basic implementation of Aggregater.
type Aggregate struct {
	// ID is the unique identifier of this aggregate.
	ID vo.UUID `json:"-"`
	// Type is the name of the aggregate.
	Type vo.AggregateType `json:"-"`
	// OrgID is the organization belongs to.
	OrgID vo.UUID `json:"-"`
	// InstanceID is the instance this aggregate belongs to.
	InstanceID vo.UUID `json:"-"`
	// Version is the semver this aggregate represents.
	Version vo.Version `json:"-"`
}

func (a Aggregate) Validate() error {
	if a.ID == "" {
		return errors.New("aggregate id must not be empty")
	}
	if a.Type == "" {
		return errors.New("aggregate type must not be empty")
	}
	if a.Version == "" {
		return errors.New("aggregate version must not be empty")
	}
	return nil
}

func (a Aggregate) IsAggregateTypes(types ...vo.AggregateType) bool {
	for _, typ := range types {
		if a.Type == typ {
			return true
		}
	}
	return false
}

func (a Aggregate) IsAggregateIDs(ids ...vo.UUID) bool {
	for _, id := range ids {
		if a.ID == id {
			return true
		}
	}
	return false
}

type aggregateOpt func(*Aggregate)

// NewAggregate is the default constructor of an aggregate
// opts overwrite values calculated by given parameters.
func NewAggregate(
	id vo.UUID,
	typ vo.AggregateType,
	version vo.Version,
	opts ...aggregateOpt,
) *Aggregate {
	a := &Aggregate{
		ID:      id,
		Type:    typ,
		Version: version,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

// WithOrgID the org ID of the aggregate option.
func WithOrgID(orgID vo.UUID) aggregateOpt {
	return func(aggregate *Aggregate) {
		aggregate.OrgID = orgID
	}
}

// WithInstanceID the instance ID of the aggregate option.
func WithInstanceID(instanceID vo.UUID) aggregateOpt {
	return func(aggregate *Aggregate) {
		aggregate.InstanceID = instanceID
	}
}
