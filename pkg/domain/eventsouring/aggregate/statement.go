package aggregate

import (
	"errors"
	"strconv"
	"strings"

	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/event"
	"github.com/galaxyobe/go-ddd/pkg/domain/eventsouring/vo"
)

var (
	ErrNoProjection = errors.New("no projection")
	ErrNoValues     = errors.New("no values")
	ErrNoCondition  = errors.New("no condition")
)

type execOption func(*execOptions)

type execOptions struct {
	tableName string
	args      []interface{}
	err       error
}

type query func(config execOptions) string

func exec(config execOptions, q query, opts []execOption) Exec {
	return func(ex vo.IExecute, projectionName string) error {
		if projectionName == "" {
			return ErrNoProjection
		}

		if config.err != nil {
			return config.err
		}

		config.tableName = projectionName
		for _, opt := range opts {
			opt(&config)
		}

		if _, err := ex.Exec(q(config), config.args...); err != nil {
			return errors.New("exec failed")
		}

		return nil
	}
}

func multiExec(execList []Exec) Exec {
	return func(ex vo.IExecute, projectionName string) error {
		for _, f := range execList {
			if err := f(ex, projectionName); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithTableSuffix(name string) func(*execOptions) {
	return func(o *execOptions) {
		o.tableName += "_" + name
	}
}

func NewCreateStatement(event event.IEvent, values []vo.Column, opts ...execOption) *vo.Statement {
	cols, params, args := columnsToQuery(values)
	columnNames := strings.Join(cols, ", ")
	valuesPlaceholder := strings.Join(params, ", ")

	options := execOptions{
		args: args,
	}

	if len(values) == 0 {
		options.err = ErrNoValues
	}

	q := func(config execOptions) string {
		return "INSERT INTO " + config.tableName + " (" + columnNames + ") VALUES (" + valuesPlaceholder + ")"
	}

	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          exec(options, q, opts),
	}
}

func NewUpsertStatement(event event.IEvent, values []vo.Column, opts ...execOption) *vo.Statement {
	cols, params, args := columnsToQuery(values)
	columnNames := strings.Join(cols, ", ")
	valuesPlaceholder := strings.Join(params, ", ")

	options := execOptions{
		args: args,
	}

	if len(values) == 0 {
		options.err = ErrNoValues
	}

	q := func(config execOptions) string {
		return "UPSERT INTO " + config.tableName + " (" + columnNames + ") VALUES (" + valuesPlaceholder + ")"
	}

	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          exec(options, q, opts),
	}
}

func NewUpdateStatement(event event.IEvent, values []vo.Column, conditions []vo.Condition, opts ...execOption) *vo.Statement {
	cols, params, args := columnsToQuery(values)
	wheres, whereArgs := conditionsToWhere(conditions, len(params))
	args = append(args, whereArgs...)

	columnNames := strings.Join(cols, ", ")
	valuesPlaceholder := strings.Join(params, ", ")
	wheresPlaceholders := strings.Join(wheres, " AND ")

	options := execOptions{
		args: args,
	}

	if len(values) == 0 {
		options.err = ErrNoValues
	}

	if len(conditions) == 0 {
		options.err = ErrNoCondition
	}

	q := func(config execOptions) string {
		return "UPDATE " + config.tableName + " SET (" + columnNames + ") = (" + valuesPlaceholder + ") WHERE " + wheresPlaceholders
	}

	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          exec(options, q, opts),
	}
}

func NewDeleteStatement(event event.IEvent, conditions []vo.Condition, opts ...execOption) *vo.Statement {
	wheres, args := conditionsToWhere(conditions, 0)

	wheresPlaceholders := strings.Join(wheres, " AND ")

	options := execOptions{
		args: args,
	}

	if len(conditions) == 0 {
		options.err = ErrNoCondition
	}

	q := func(config execOptions) string {
		return "DELETE FROM " + config.tableName + " WHERE " + wheresPlaceholders
	}

	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          exec(options, q, opts),
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

func NewMultiStatement(event event.IEvent, opts ...func(event.IEvent) Exec) *vo.Statement {
	if len(opts) == 0 {
		return NewNoOpStatement(event)
	}
	execs := make([]Exec, len(opts))
	for i, opt := range opts {
		execs[i] = opt(event)
	}

	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          multiExec(execs),
	}
}

type Exec func(ex vo.IExecute, projectionName string) error

func AddCreateStatement(columns []vo.Column, opts ...execOption) func(event.IEvent) Exec {
	return func(event event.IEvent) Exec {
		return NewCreateStatement(event, columns, opts...).Execute
	}
}

func AddUpsertStatement(values []vo.Column, opts ...execOption) func(event.IEvent) Exec {
	return func(event event.IEvent) Exec {
		return NewUpsertStatement(event, values, opts...).Execute
	}
}

func AddUpdateStatement(values []vo.Column, conditions []vo.Condition, opts ...execOption) func(event.IEvent) Exec {
	return func(event event.IEvent) Exec {
		return NewUpdateStatement(event, values, conditions, opts...).Execute
	}
}

func AddDeleteStatement(conditions []vo.Condition, opts ...execOption) func(event.IEvent) Exec {
	return func(event event.IEvent) Exec {
		return NewDeleteStatement(event, conditions, opts...).Execute
	}
}

// NewCopyStatement creates a new upsert statement which updates a column from an existing row
// cols represent the columns which are objective to change.
// if the value of a col is empty the data will be copied from the selected row
// if the value of a col is not empty the data will be set by the static value
// conditions represent the conditions for the selection subquery
func NewCopyStatement(event event.IEvent, columns []vo.Column, conditions []vo.Condition, opts ...execOption) *vo.Statement {
	columnNames := make([]string, len(columns))
	selectColumns := make([]string, len(columns))
	argCounter := 0
	args := []interface{}{}

	for i, col := range columns {
		columnNames[i] = col.Name
		selectColumns[i] = col.Name
		if col.Value != nil {
			argCounter++
			selectColumns[i] = "$" + strconv.Itoa(argCounter)
			args = append(args, col.Value)
		}
	}

	wheres := make([]string, len(conditions))
	for i, cond := range conditions {
		argCounter++
		wheres[i] = "copy_table." + cond.Name + " = $" + strconv.Itoa(argCounter)
		args = append(args, cond.Value)
	}

	config := execOptions{
		args: args,
	}

	if len(columns) == 0 {
		config.err = ErrNoValues
	}

	if len(conditions) == 0 {
		config.err = ErrNoCondition
	}

	q := func(config execOptions) string {
		return "UPSERT INTO " +
			config.tableName +
			" (" +
			strings.Join(columnNames, ", ") +
			") SELECT " +
			strings.Join(selectColumns, ", ") +
			" FROM " +
			config.tableName + " AS copy_table WHERE " +
			strings.Join(wheres, " AND ")
	}

	aggregate := event.GetAggregate()
	return &vo.Statement{
		AggregateType:    aggregate.Type,
		Sequence:         event.GetSequence(),
		PreviousSequence: event.GetPreviousAggregateTypeSequence(),
		InstanceID:       aggregate.InstanceID,
		Execute:          exec(config, q, opts),
	}
}

func columnsToQuery(cols []vo.Column) (names []string, parameters []string, values []interface{}) {
	names = make([]string, len(cols))
	values = make([]interface{}, len(cols))
	parameters = make([]string, len(cols))
	for i, col := range cols {
		names[i] = col.Name
		values[i] = col.Value
		parameters[i] = "$" + strconv.Itoa(i+1)
		if col.Option != nil {
			parameters[i] = col.Option(parameters[i])
		}
	}
	return names, parameters, values
}

func conditionsToWhere(cols []vo.Condition, paramOffset int) (wheres []string, values []interface{}) {
	wheres = make([]string, len(cols))
	values = make([]interface{}, len(cols))

	for i, col := range cols {
		wheres[i] = "(" + col.Name + " = $" + strconv.Itoa(i+1+paramOffset) + ")"
		values[i] = col.Value
	}

	return wheres, values
}
