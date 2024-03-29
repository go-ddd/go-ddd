// Code generated by ent, DO NOT EDIT.

package do

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/galaxyobe/go-ddd/pkg/domain/database/do"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/repository/do/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
	"github.com/galaxyobe/go-ddd/pkg/types"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
}

// SetAggregateID sets the "aggregate_id" field.
func (ec *EventCreate) SetAggregateID(t types.UUID) *EventCreate {
	ec.mutation.SetAggregateID(t)
	return ec
}

// SetOrgID sets the "org_id" field.
func (ec *EventCreate) SetOrgID(t types.UUID) *EventCreate {
	ec.mutation.SetOrgID(t)
	return ec
}

// SetInstanceID sets the "instance_id" field.
func (ec *EventCreate) SetInstanceID(t types.UUID) *EventCreate {
	ec.mutation.SetInstanceID(t)
	return ec
}

// SetVersion sets the "version" field.
func (ec *EventCreate) SetVersion(t types.Version) *EventCreate {
	ec.mutation.SetVersion(t)
	return ec
}

// SetCreator sets the "creator" field.
func (ec *EventCreate) SetCreator(t types.UUID) *EventCreate {
	ec.mutation.SetCreator(t)
	return ec
}

// SetType sets the "type" field.
func (ec *EventCreate) SetType(vt vo.EventType) *EventCreate {
	ec.mutation.SetType(vt)
	return ec
}

// SetAggregateType sets the "aggregate_type" field.
func (ec *EventCreate) SetAggregateType(vt vo.AggregateType) *EventCreate {
	ec.mutation.SetAggregateType(vt)
	return ec
}

// SetMetadata sets the "metadata" field.
func (ec *EventCreate) SetMetadata(dm do.StringMap) *EventCreate {
	ec.mutation.SetMetadata(dm)
	return ec
}

// SetData sets the "data" field.
func (ec *EventCreate) SetData(b []byte) *EventCreate {
	ec.mutation.SetData(b)
	return ec
}

// SetSequence sets the "sequence" field.
func (ec *EventCreate) SetSequence(u uint64) *EventCreate {
	ec.mutation.SetSequence(u)
	return ec
}

// SetPreviousAggregateSequence sets the "previous_aggregate_sequence" field.
func (ec *EventCreate) SetPreviousAggregateSequence(u uint64) *EventCreate {
	ec.mutation.SetPreviousAggregateSequence(u)
	return ec
}

// SetPreviousAggregateTypeSequence sets the "previous_aggregate_type_sequence" field.
func (ec *EventCreate) SetPreviousAggregateTypeSequence(u uint64) *EventCreate {
	ec.mutation.SetPreviousAggregateTypeSequence(u)
	return ec
}

// SetService sets the "service" field.
func (ec *EventCreate) SetService(s string) *EventCreate {
	ec.mutation.SetService(s)
	return ec
}

// SetCreateTime sets the "create_time" field.
func (ec *EventCreate) SetCreateTime(t time.Time) *EventCreate {
	ec.mutation.SetCreateTime(t)
	return ec
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ec *EventCreate) SetNillableCreateTime(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetCreateTime(*t)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EventCreate) SetID(t types.UUID) *EventCreate {
	ec.mutation.SetID(t)
	return ec
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("do: uninitialized hook (forgotten import do/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Event)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() {
	if _, ok := ec.mutation.CreateTime(); !ok {
		v := event.DefaultCreateTime
		ec.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.AggregateID(); !ok {
		return &ValidationError{Name: "aggregate_id", err: errors.New(`do: missing required field "Event.aggregate_id"`)}
	}
	if v, ok := ec.mutation.AggregateID(); ok {
		if err := event.AggregateIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "aggregate_id", err: fmt.Errorf(`do: validator failed for field "Event.aggregate_id": %w`, err)}
		}
	}
	if _, ok := ec.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`do: missing required field "Event.org_id"`)}
	}
	if v, ok := ec.mutation.OrgID(); ok {
		if err := event.OrgIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "org_id", err: fmt.Errorf(`do: validator failed for field "Event.org_id": %w`, err)}
		}
	}
	if _, ok := ec.mutation.InstanceID(); !ok {
		return &ValidationError{Name: "instance_id", err: errors.New(`do: missing required field "Event.instance_id"`)}
	}
	if v, ok := ec.mutation.InstanceID(); ok {
		if err := event.InstanceIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "instance_id", err: fmt.Errorf(`do: validator failed for field "Event.instance_id": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`do: missing required field "Event.version"`)}
	}
	if v, ok := ec.mutation.Version(); ok {
		if err := event.VersionValidator(string(v)); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`do: validator failed for field "Event.version": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Creator(); !ok {
		return &ValidationError{Name: "creator", err: errors.New(`do: missing required field "Event.creator"`)}
	}
	if v, ok := ec.mutation.Creator(); ok {
		if err := event.CreatorValidator(string(v)); err != nil {
			return &ValidationError{Name: "creator", err: fmt.Errorf(`do: validator failed for field "Event.creator": %w`, err)}
		}
	}
	if _, ok := ec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`do: missing required field "Event.type"`)}
	}
	if v, ok := ec.mutation.GetType(); ok {
		if err := event.TypeValidator(string(v)); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`do: validator failed for field "Event.type": %w`, err)}
		}
	}
	if _, ok := ec.mutation.AggregateType(); !ok {
		return &ValidationError{Name: "aggregate_type", err: errors.New(`do: missing required field "Event.aggregate_type"`)}
	}
	if v, ok := ec.mutation.AggregateType(); ok {
		if err := event.AggregateTypeValidator(string(v)); err != nil {
			return &ValidationError{Name: "aggregate_type", err: fmt.Errorf(`do: validator failed for field "Event.aggregate_type": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Sequence(); !ok {
		return &ValidationError{Name: "sequence", err: errors.New(`do: missing required field "Event.sequence"`)}
	}
	if _, ok := ec.mutation.PreviousAggregateSequence(); !ok {
		return &ValidationError{Name: "previous_aggregate_sequence", err: errors.New(`do: missing required field "Event.previous_aggregate_sequence"`)}
	}
	if _, ok := ec.mutation.PreviousAggregateTypeSequence(); !ok {
		return &ValidationError{Name: "previous_aggregate_type_sequence", err: errors.New(`do: missing required field "Event.previous_aggregate_type_sequence"`)}
	}
	if _, ok := ec.mutation.Service(); !ok {
		return &ValidationError{Name: "service", err: errors.New(`do: missing required field "Event.service"`)}
	}
	if v, ok := ec.mutation.Service(); ok {
		if err := event.ServiceValidator(v); err != nil {
			return &ValidationError{Name: "service", err: fmt.Errorf(`do: validator failed for field "Event.service": %w`, err)}
		}
	}
	if _, ok := ec.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`do: missing required field "Event.create_time"`)}
	}
	if v, ok := ec.mutation.ID(); ok {
		if err := event.IDValidator(string(v)); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`do: validator failed for field "Event.id": %w`, err)}
		}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(types.UUID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Event.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: event.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: event.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.AggregateID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldAggregateID,
		})
		_node.AggregateID = value
	}
	if value, ok := ec.mutation.OrgID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldOrgID,
		})
		_node.OrgID = value
	}
	if value, ok := ec.mutation.InstanceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldInstanceID,
		})
		_node.InstanceID = value
	}
	if value, ok := ec.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldVersion,
		})
		_node.Version = value
	}
	if value, ok := ec.mutation.Creator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldCreator,
		})
		_node.Creator = value
	}
	if value, ok := ec.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldType,
		})
		_node.Type = value
	}
	if value, ok := ec.mutation.AggregateType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldAggregateType,
		})
		_node.AggregateType = value
	}
	if value, ok := ec.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: event.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := ec.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: event.FieldData,
		})
		_node.Data = value
	}
	if value, ok := ec.mutation.Sequence(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: event.FieldSequence,
		})
		_node.Sequence = value
	}
	if value, ok := ec.mutation.PreviousAggregateSequence(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: event.FieldPreviousAggregateSequence,
		})
		_node.PreviousAggregateSequence = value
	}
	if value, ok := ec.mutation.PreviousAggregateTypeSequence(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: event.FieldPreviousAggregateTypeSequence,
		})
		_node.PreviousAggregateTypeSequence = value
	}
	if value, ok := ec.mutation.Service(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldService,
		})
		_node.Service = value
	}
	if value, ok := ec.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	return _node, _spec
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	builders []*EventCreate
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
