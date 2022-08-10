package vo

type ExecOption func(*ExecOptions)

type ExecOptions struct {
	tableName string
	args      []interface{}
	err       error
}

type Query func(options ExecOptions) string
