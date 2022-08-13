package interfaces

import (
	"context"
	"database/sql"

	"github.com/galaxyobe/go-ddd/pkg/domain/database/vo"
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

// IPrepare is the interface that wraps the Prepare method.
//
// Prepare executes the given query as implemented by database/sql.Prepare.
type IPrepare interface {
	Prepare(query string) (*sql.Stmt, error)
}

// IPrepareContext is the interface that wraps the Prepare and PrepareContext methods.
//
// Prepare executes the given query as implemented by database/sql.Prepare.
// PrepareContext executes the given query as implemented by database/sql.PrepareContext.
type IPrepareContext interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type IStatement interface {
	Create(*vo.ExecOptions) (query string, args []any)
	Upsert(*vo.ExecOptions) (query string, args []any)
	Update(*vo.ExecOptions) (query string, args []any)
	Delete(*vo.ExecOptions) (query string, args []any)
}
