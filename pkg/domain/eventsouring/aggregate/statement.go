package aggregate

import (
	"errors"

	"github.com/galaxyobe/go-ddd/pkg/domain/database/interfaces"
	dbvo "github.com/galaxyobe/go-ddd/pkg/domain/database/vo"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

var (
	ErrNoValues    = errors.New("no values")
	ErrNoCondition = errors.New("no condition")
)

type ProjectionStatement struct {
	interfaces.IStatement
}

func (p ProjectionStatement) NewCreateStatement(event event.IEvent, columns []dbvo.Column, opts ...dbvo.ExecOption) *vo.Statement {
	options := dbvo.ExecOptions{
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

func (p ProjectionStatement) NewUpsertStatement(event event.IEvent, columns []dbvo.Column, opts ...dbvo.ExecOption) *vo.Statement {
	options := dbvo.ExecOptions{
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

//
//
// func NewUpdateStatement(event event.IEvent, values []dbvo.Column, conditions []dbvo.Condition, opts ...dbvo.ExecOption) *vo.Statement {
// 	options := dbvo.ExecOptions{
// 		Columns: columns,
// 	}
// 	aggregate := event.GetAggregate()
// 	return &vo.Statement{
// 		AggregateType:    aggregate.Type,
// 		Sequence:         event.GetSequence(),
// 		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
// 		InstanceID:       aggregate.InstanceID,
// 		Execute:          vo.ExecProjection(options, p.Upsert, opts),
// 	}
// }
//
// func NewDeleteStatement(event event.IEvent, conditions []dbvo.Condition, opts ...dbvo.ExecOption) *vo.Statement {
// 	wheres, args := conditionsToWhere(conditions, 0)
//
// 	wheresPlaceholders := strings.Join(wheres, " AND ")
//
// 	options := execOptions{
// 		args: args,
// 	}
//
// 	if len(conditions) == 0 {
// 		options.err = ErrNoCondition
// 	}
//
// 	q := func(config execOptions) string {
// 		return "DELETE FROM " + config.tableName + " WHERE " + wheresPlaceholders
// 	}
//
// 	aggregate := event.GetAggregate()
// 	return &vo.Statement{
// 		AggregateType:    aggregate.Type,
// 		Sequence:         event.GetSequence(),
// 		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
// 		InstanceID:       aggregate.InstanceID,
// 		Execute:          exec(options, q, opts),
// 	}
// }

func NewNoOpStatement(event event.IEvent) *vo.Statement {
	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
	}
}

// func NewMultiStatement(event event.IEvent, opts ...func(event.IEvent) Exec) *vo.Statement {
// 	if len(opts) == 0 {
// 		return NewNoOpStatement(event)
// 	}
// 	execs := make([]Exec, len(opts))
// 	for i, opt := range opts {
// 		execs[i] = opt(event)
// 	}
//
// 	aggregate := event.GetAggregate()
// 	return &vo.Statement{
// 		AggregateType:    aggregate.Type,
// 		Sequence:         event.GetSequence(),
// 		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
// 		InstanceID:       aggregate.InstanceID,
// 		Execute:          multiExec(execs),
// 	}
// }
//
// func AddCreateStatement(columns []vo.Column, opts ...execOption) func(event.IEvent) Exec {
// 	return func(event event.IEvent) Exec {
// 		return NewCreateStatement(event, columns, opts...).Execute
// 	}
// }
//
// func AddUpsertStatement(columns []dbvo.Column, opts ...dbvo.ExecOption) func(event.IEvent) Exec {
// 	return func(event event.IEvent) Exec {
// 		return NewUpsertStatement(event, values, opts...).Execute
// 	}
// }
//
// func AddUpdateStatement(values []vo.Column, conditions []vo.Condition, opts ...execOption) func(event.IEvent) Exec {
// 	return func(event event.IEvent) Exec {
// 		return NewUpdateStatement(event, values, conditions, opts...).Execute
// 	}
// }
//
// func AddDeleteStatement(conditions []vo.Condition, opts ...execOption) func(event.IEvent) Exec {
// 	return func(event event.IEvent) Exec {
// 		return NewDeleteStatement(event, conditions, opts...).Execute
// 	}
// }
//
// // NewCopyStatement creates a new upsert statement which updates a column from an existing row
// // cols represent the columns which are objective to change.
// // if the value of a col is empty the data will be copied from the selected row
// // if the value of a col is not empty the data will be set by the static value
// // conditions represent the conditions for the selection subquery
// func NewCopyStatement(event event.IEvent, columns []vo.Column, conditions []vo.Condition, opts ...execOption) *vo.Statement {
// 	columnNames := make([]string, len(columns))
// 	selectColumns := make([]string, len(columns))
// 	argCounter := 0
// 	args := []interface{}{}
//
// 	for i, col := range columns {
// 		columnNames[i] = col.Name
// 		selectColumns[i] = col.Name
// 		if col.Value != nil {
// 			argCounter++
// 			selectColumns[i] = "$" + strconv.Itoa(argCounter)
// 			args = append(args, col.Value)
// 		}
// 	}
//
// 	wheres := make([]string, len(conditions))
// 	for i, cond := range conditions {
// 		argCounter++
// 		wheres[i] = "copy_table." + cond.Name + " = $" + strconv.Itoa(argCounter)
// 		args = append(args, cond.Value)
// 	}
//
// 	config := execOptions{
// 		args: args,
// 	}
//
// 	if len(columns) == 0 {
// 		config.err = ErrNoValues
// 	}
//
// 	if len(conditions) == 0 {
// 		config.err = ErrNoCondition
// 	}
//
// 	q := func(config execOptions) string {
// 		return "UPSERT INTO " +
// 			config.tableName +
// 			" (" +
// 			strings.Join(columnNames, ", ") +
// 			") SELECT " +
// 			strings.Join(selectColumns, ", ") +
// 			" FROM " +
// 			config.tableName + " AS copy_table WHERE " +
// 			strings.Join(wheres, " AND ")
// 	}
//
// 	aggregate := event.GetAggregate()
// 	return &vo.Statement{
// 		AggregateType:    aggregate.Type,
// 		Sequence:         event.GetSequence(),
// 		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
// 		InstanceID:       aggregate.InstanceID,
// 		Execute:          exec(config, q, opts),
// 	}
// }
