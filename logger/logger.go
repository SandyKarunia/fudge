package logger

import (
	"github.com/sandykarunia/fudge/sdk"
	"os"
)

// Logger logs messages. It only includes the necessary severity level which Fudge uses.
//go:generate mockery -name=Logger
type Logger interface {
	Info(tag string, message string, args ...interface{})
	Warn(tag string, message string, args ...interface{})
	Error(tag string, message string, args ...interface{})
}

// stdLogger logs messages to stdout or stderr.
type stdLogger struct {
	fmt sdk.FmtFunctions
}

func (l *stdLogger) Info(tag string, message string, args ...interface{}) {
	l.doLog(os.Stdout, "INFO", tag, message, args...)
}

func (l *stdLogger) Warn(tag string, message string, args ...interface{}) {
	l.doLog(os.Stdout, "WARN", tag, message, args...)
}

func (l *stdLogger) Error(tag string, message string, args ...interface{}) {
	l.doLog(os.Stderr, "ERROR", tag, message, args...)
}

func (l *stdLogger) doLog(
	targetFile *os.File,
	severityTag string,
	tag string,
	message string,
	args ...interface{},
) {
	_, _ = l.fmt.Fprintf(targetFile, "["+severityTag+"] ["+tag+"] "+message+"\n", args...)
}
