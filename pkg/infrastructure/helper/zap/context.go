package zap_helper

import (
	"context"

	"go.uber.org/zap"
)

type (
	loggerCtx struct{}
	fieldCtx  struct{}
)

func WithLogger(ctx context.Context, logger ILogger) context.Context {
	if logger == nil {
		return ctx
	}
	return context.WithValue(ctx, loggerCtx{}, logger)
}

func NewLoggerFromContext(ctx context.Context) ILogger {
	logger, _ := ctx.Value(loggerCtx{}).(ILogger)
	if logger == nil {
		return Wrap(zap.NewNop())
	}
	return logger.WithContext(ctx)
}

func WithField(ctx context.Context, fields ...zap.Field) context.Context {
	if len(fields) == 0 {
		return ctx
	}
	return context.WithValue(ctx, fieldCtx{}, fields)
}

func GetFieldsFromContext(ctx context.Context) []zap.Field {
	fields, _ := ctx.Value(fieldCtx{}).([]zap.Field)
	return fields
}
