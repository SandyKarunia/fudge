package logger

import (
	utilsmocks "github.com/sandykarunia/fudge/utils/mocks"
	"testing"
	"time"
)

func TestLoggerImpl_NoCrash(t *testing.T) {
	sysMocks := &utilsmocks.System{}
	sysMocks.On("GetFudgeDir").Return("")
	obj := (&loggerImpl{
		sys: sysMocks,
	}).init()

	obj.Info("info message no arguments")
	obj.Info("info message %d", time.Now().Unix())
	obj.Warn("warn message no arguments")
	obj.Warn("warn message %d", time.Now().Unix())
	obj.Error("error message no arguments")
	obj.Error("error message %d", time.Now().Unix())
}
