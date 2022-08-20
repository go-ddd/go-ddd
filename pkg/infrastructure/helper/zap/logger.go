package zap_helper

import (
	"context"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(logger *zap.Logger) ILogger {
	return &Logger{
		Logger: logger,
	}
}

func (log *Logger) WithContext(ctx context.Context) ILogger {
	fields := GetFieldsFromContext(ctx)
	if len(fields) == 0 {
		return log
	}
	return &Logger{Logger: log.With(fields...)}
}

func (log *Logger) OnError(err error, fields ...zap.Field) IZapLog {
	switch {
	case err == nil:
		return NewNopOut()
	case len(fields) == 0:
		return &Logger{Logger: log.With(zap.Error(err))}
	default:
		return &Logger{Logger: log.With(append(fields, zap.Error(err))...)}
	}
}
