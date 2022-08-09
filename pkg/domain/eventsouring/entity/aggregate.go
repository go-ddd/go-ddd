package entity

import (
	"errors"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
	"github.com/galaxyobe/go-ddd/pkg/infrastructure/tools"
)

// Aggregate is the basic implementation of Aggregater.
type Aggregate struct {
	// ID is the unique identifier of this aggregate.
	ID vo.GUID `json:"-"`
	// Type is the name of the aggregate.
	Type vo.AggregateType `json:"-"`
	// OrgID is the organization belongs to.
	OrgID string `json:"-"`
	// InstanceID is the instance this aggregate belongs to.
	InstanceID string `json:"-"`
	// Version is the semver this aggregate represents.
	Version vo.Version `json:"-"`
}

func (a Aggregate) Validate() error {
	if tools.IsGUIDNull(a.ID) {
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

func (a Aggregate) IsAggregateIDs(ids ...vo.GUID) bool {
	for _, id := range ids {
		if a.ID.Equaled(id) {
			return true
		}
	}
	return false
}

type aggregateOpt func(*Aggregate)

// NewAggregate is the default constructor of an aggregate
// opts overwrite values calculated by given parameters.
func NewAggregate(
	id vo.GUID,
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
func WithOrgID(orgID string) aggregateOpt {
	return func(aggregate *Aggregate) {
		aggregate.OrgID = orgID
	}
}

// WithInstanceID the instance ID of the aggregate option.
func WithInstanceID(instanceID string) aggregateOpt {
	return func(aggregate *Aggregate) {
		aggregate.InstanceID = instanceID
	}
}
