package logger

import (
	"testing"
	"time"
)

func TestLoggerImpl_NoCrash(t *testing.T) {
	obj := &loggerImpl{}
	obj.Info("info message no arguments")
	obj.Info("info message %d", time.Now().Unix())
	obj.Warn("warn message no arguments")
	obj.Warn("warn message %d", time.Now().Unix())
	obj.Error("error message no arguments")
	obj.Error("error message %d", time.Now().Unix())
}
