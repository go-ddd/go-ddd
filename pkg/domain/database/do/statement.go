package do

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ExecOption func(*ExecOptions)

type ExecOptions struct {
	TableName  string
	Columns    []Column
	Conditions []Condition
	Err        error
}

func WithTableSuffix(name string) ExecOption {
	return func(o *ExecOptions) {
		o.TableName += "_" + name
	}
}

type Column struct {
	Name   string
	Value  interface{}
	Option func(string) string
	Err    error
}

func NewCol(name string, value interface{}) Column {
	return Column{
		Name:  name,
		Value: value,
	}
}

func NewJSONCol(name string, value interface{}) Column {
	marshalled, err := json.Marshal(value)
	return Column{
		Name:  name,
		Value: marshalled,
		Err:   err,
	}
}

type Condition Column

func NewCond(name string, value interface{}) Condition {
	return Condition{
		Name:  name,
		Value: value,
	}
}

func ColumnsToQuery(cols []Column) ([]string, []string, []interface{}, error) {
	var (
		names      = make([]string, len(cols))
		values     = make([]interface{}, len(cols))
		parameters = make([]string, len(cols))
	)
	for i, col := range cols {
		if col.Err != nil {
			return nil, nil, nil, fmt.Errorf("%s: %w", col.Name, col.Err)
		}
		names[i] = col.Name
		values[i] = col.Value
		parameters[i] = "$" + strconv.Itoa(i+1)
		if col.Option != nil {
			parameters[i] = col.Option(parameters[i])
		}
	}
	return names, parameters, values, nil
}

func ConditionsToWhere(cols []Condition, paramOffset int) (wheres []string, values []interface{}) {
	wheres = make([]string, len(cols))
	values = make([]interface{}, len(cols))

	for i, col := range cols {
		wheres[i] = "(" + col.Name + " = $" + strconv.Itoa(i+1+paramOffset) + ")"
		values[i] = col.Value
	}

	return wheres, values
}
