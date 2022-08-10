package interfaces

import (
	"context"
	"database/sql"
)

// IExec is the interface that wraps the Exec method.
//
// Exec executes the given query as implemented by database/sql.Exec.
type IExec interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// IQuery is the interface that wraps the Query method.
//
// Query executes the given query as implemented by database/sql.Query.
type IQuery interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// IExecContext is the interface that wraps the ExecContext method.
//
// Exec executes the given query as implemented by database/sql.ExecContext.
type IExecContext interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// IQueryContext is the interface that wraps the QueryContext method.
//
// QueryContext executes the given query as implemented by database/sql.QueryContext.
type IQueryContext interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}
