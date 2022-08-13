package projection

import (
	"github.com/galaxyobe/go-ddd/pkg/domain/database/interfaces"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

type Statement struct {
	interfaces.IStatement
}

func (p Statement) NewCreateStatement(event event.IEvent, columns []vo.Column, opts ...vo.ExecOption) *vo.Statement {
	options := vo.ExecOptions{
		Columns: columns,
	}
	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          vo.ExecProjection(options, p.Create, opts),
	}
}

func (p Statement) NewUpsertStatement(event event.IEvent, columns []vo.Column, opts ...vo.ExecOption) *vo.Statement {
	options := vo.ExecOptions{
		Columns: columns,
	}
	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          vo.ExecProjection(options, p.Upsert, opts),
	}
}

func (p Statement) NewUpdateStatement(event event.IEvent, columns []vo.Column, conditions []vo.Condition, opts ...vo.ExecOption) *vo.Statement {
	options := vo.ExecOptions{
		Columns:    columns,
		Conditions: conditions,
	}
	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          vo.ExecProjection(options, p.Upsert, opts),
	}
}

func (p Statement) NewDeleteStatement(event event.IEvent, conditions []vo.Condition, opts ...vo.ExecOption) *vo.Statement {
	options := vo.ExecOptions{
		Conditions: conditions,
	}
	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          vo.ExecProjection(options, p.Delete, opts),
	}
}

func (p Statement) AddCreateStatement(columns []vo.Column, opts ...vo.ExecOption) func(event.IEvent) vo.ProjectionExecute {
	return func(event event.IEvent) vo.ProjectionExecute {
		return p.NewCreateStatement(event, columns, opts...).Execute
	}
}

func (p Statement) AddUpsertStatement(columns []vo.Column, opts ...vo.ExecOption) func(event.IEvent) vo.ProjectionExecute {
	return func(event event.IEvent) vo.ProjectionExecute {
		return p.NewUpsertStatement(event, columns, opts...).Execute
	}
}

func (p Statement) AddUpdateStatement(columns []vo.Column, conditions []vo.Condition, opts ...vo.ExecOption) func(event.IEvent) vo.ProjectionExecute {
	return func(event event.IEvent) vo.ProjectionExecute {
		return p.NewUpdateStatement(event, columns, conditions, opts...).Execute
	}
}

func (p Statement) AddDeleteStatement(conditions []vo.Condition, opts ...vo.ExecOption) func(event.IEvent) vo.ProjectionExecute {
	return func(event event.IEvent) vo.ProjectionExecute {
		return p.NewDeleteStatement(event, conditions, opts...).Execute
	}
}

func NewNoOpStatement(event event.IEvent) *vo.Statement {
	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
	}
}

func NewMultiStatement(event event.IEvent, opts ...func(event.IEvent) vo.ProjectionExecute) *vo.Statement {
	if len(opts) == 0 {
		return NewNoOpStatement(event)
	}
	execs := make([]vo.ProjectionExecute, len(opts))
	for i, opt := range opts {
		execs[i] = opt(event)
	}
	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          vo.MultiExecProjection(execs),
	}
}
