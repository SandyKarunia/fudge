package logger

import (
	"gitlab.com/sandykarunia/fudge/sdk/mocks"
	"os"
	"testing"
)

func setupTest(targetFile *os.File, severityTag, tag, msg string, arg interface{}) *stdLogger {
	mockFmt := &mocks.FmtFunctions{}
	obj := &stdLogger{fmt: mockFmt}
	mockFmt.On("Fprintf", targetFile, "["+severityTag+"] ["+tag+"] "+msg+"\n", arg).
		Return(123, nil).
		Once()
	return obj
}

func verifyTest(t *testing.T, l *stdLogger) {
	mf, ok := l.fmt.(*mocks.FmtFunctions)
	if !ok {
		t.Fatalf("cannot cast l.fmt to *mocks.FmtFunctions")
	}
	mf.AssertExpectations(t)
}

func TestStdLogger_Info(t *testing.T) {
	obj := setupTest(os.Stdout, "INFO", "123", "msg", 1)
	obj.Info("123", "msg", 1)
	verifyTest(t, obj)
}

func TestStdLogger_Warn(t *testing.T) {
	obj := setupTest(os.Stdout, "WARN", "123", "msg", 1)
	obj.Warn("123", "msg", 1)
	verifyTest(t, obj)
}

func TestStdLogger_Error(t *testing.T) {
	obj := setupTest(os.Stderr, "ERROR", "123", "msg", 1)
	obj.Error("123", "msg", 1)
	verifyTest(t, obj)
}
