package zap_helper

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type IZapLog interface {
	Log(lvl zapcore.Level, msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

type IZapLogger interface {
	IZapLog
	Sugar() *zap.SugaredLogger
	Named(s string) *zap.Logger
	WithOptions(opts ...zap.Option) *zap.Logger
	With(fields ...zap.Field) *zap.Logger
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
	Sync() error
	Core() zapcore.Core
}

type ILogger interface {
	IZapLogger
	WithContext(ctx context.Context) ILogger
	OnError(err error, fields ...zap.Field) IZapLog
}
