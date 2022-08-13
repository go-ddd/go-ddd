package zap_helper

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		Logger: logger,
	}
}

func (log *Logger) OnError(err error, fields ...zap.Field) *Logger {
	switch {
	case err == nil:
		return &Logger{Logger: zap.NewNop()}
	case len(fields) == 0:
		return &Logger{Logger: log.With(zap.Error(err))}
	default:
		return &Logger{Logger: log.With(append(fields, zap.Error(err))...)}
	}
}
