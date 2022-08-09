package vo

import (
	"database/sql"
)

type Statements []Statement

func (stmts Statements) Len() int           { return len(stmts) }
func (stmts Statements) Swap(i, j int)      { stmts[i], stmts[j] = stmts[j], stmts[i] }
func (stmts Statements) Less(i, j int) bool { return stmts[i].Sequence < stmts[j].Sequence }

type IExecute interface {
	Exec(string, ...interface{}) (sql.Result, error)
}

type Statement struct {
	AggregateType    AggregateType
	Sequence         uint64
	PreviousSequence uint64
	InstanceID       string

	Execute func(ex IExecute, projectionName string) error
}

func (s *Statement) IsNoop() bool {
	return s.Execute == nil
}

type Column struct {
	Name   string
	Value  interface{}
	Option func(string) string
}

func NewCol(name string, value interface{}) Column {
	return Column{
		Name:  name,
		Value: value,
	}
}

type Condition Column

func NewCond(name string, value interface{}) Condition {
	return Condition{
		Name:  name,
		Value: value,
	}
}
