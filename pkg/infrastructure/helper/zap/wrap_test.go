package zap_helper

import (
	"fmt"
	"testing"
	"unsafe"

	"go.uber.org/zap"
)

func TestOnError(t *testing.T) {
	t.Log("Zap Logger Sizeof:", unsafe.Sizeof(zap.Logger{}))
	t.Log("Logger Sizeof:", unsafe.Sizeof(Logger{}))
	t.Log("Wrap Logger Sizeof:", unsafe.Sizeof(wrapLogger{}))
	zapLogger, _ := zap.NewDevelopment(zap.AddStacktrace(zap.FatalLevel + 1))
	Wrap(zapLogger).OnError(nil).Error("no out")
	Wrap(zapLogger).OnError(fmt.Errorf("test error")).Error("need out")
}
