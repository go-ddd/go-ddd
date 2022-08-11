package vo

import (
	"context"
	"errors"

	"github.com/galaxyobe/go-ddd/pkg/domain/database/interfaces"
	"github.com/galaxyobe/go-ddd/pkg/domain/database/vo"
)

type Statements []Statement

func (stmts Statements) Len() int           { return len(stmts) }
func (stmts Statements) Swap(i, j int)      { stmts[i], stmts[j] = stmts[j], stmts[i] }
func (stmts Statements) Less(i, j int) bool { return stmts[i].Sequence < stmts[j].Sequence }

type ProjectionExecute func(ctx context.Context, ex interfaces.IExecContext, projectionName string) error
type Query func(*vo.ExecOptions) (string, []any)

type Statement struct {
	AggregateType    AggregateType
	Sequence         uint64
	PreviousSequence uint64
	InstanceID       string

	Execute ProjectionExecute
}

func (s *Statement) IsNoop() bool {
	return s.Execute == nil
}

var (
	ErrNoProjection = errors.New("no projection")
	ErrExecFailed   = errors.New("exec failed")
)

func ExecProjection(options vo.ExecOptions, q Query, opts []vo.ExecOption) ProjectionExecute {
	return func(ctx context.Context, ex interfaces.IExecContext, projectionName string) error {
		if projectionName == "" {
			return ErrNoProjection
		}

		if options.Err != nil {
			return options.Err
		}
		options.TableName = projectionName
		for _, opt := range opts {
			opt(&options)
		}

		query, args := q(&options)
		if options.Err != nil {
			return options.Err
		}

		if _, err := ex.ExecContext(ctx, query, args...); err != nil {
			return ErrExecFailed
		}

		return nil
	}
}

func MultiExecProjection(execs []ProjectionExecute) ProjectionExecute {
	return func(ctx context.Context, ex interfaces.IExecContext, projectionName string) error {
		for _, f := range execs {
			if err := f(ctx, ex, projectionName); err != nil {
				return err
			}
		}
		return nil
	}
}
