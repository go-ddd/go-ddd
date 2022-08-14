package zap_helper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type nopOut struct{}

func NewNopOut() IZapLog {
	return nopOut{}
}

func (n nopOut) Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	return
}

func (n nopOut) Debug(msg string, fields ...zap.Field) {
	return
}

func (n nopOut) Info(msg string, fields ...zap.Field) {
	return
}

func (n nopOut) Warn(msg string, fields ...zap.Field) {
	return
}

func (n nopOut) Error(msg string, fields ...zap.Field) {
	return
}

func (n nopOut) DPanic(msg string, fields ...zap.Field) {
	return
}

func (n nopOut) Panic(msg string, fields ...zap.Field) {
	return
}

func (n nopOut) Fatal(msg string, fields ...zap.Field) {
	return
}
