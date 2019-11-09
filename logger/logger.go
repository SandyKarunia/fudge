package logger

import (
	"fmt"
	"github.com/sandykarunia/fudge/sdk"
	"os"
	"time"
)

type severityLevel string

const (
	severityInfo  severityLevel = "INFO"
	severityWarn                = "WARN"
	severityError               = "ERROR"
)

// Logger logs messages with severity levels to stdout / stderr and a log file
// Log files will be located inside fudge directory (${HOME}/.fudge) with format log-${logger_init_timestamp}.txt
// We split the severity levels into multiple functions for convenience
//go:generate mockery -name=Logger
type Logger interface {
	// Info logs a message that acts as an information.
	// logs to stdout and log file.
	// example: "Judging submission with uuid: jcviw-rnifw-vmske-qjnd2"
	Info(message string, args ...interface{})

	// Warn logs a warning that is not necessarily fatal.
	// logs to stdout and log file.
	// example: "Fetch from input_url failed! Will retry after 1 second."
	Warn(message string, args ...interface{})

	// Error logs an error that is fatal and should not happen.
	// logs to stderr and log file.
	// example: "Failed to get user home directory information, error: $home is not defined"
	Error(message string, args ...interface{})
}

type loggerImpl struct {
	os sdk.OSFunctions
}

func (l *loggerImpl) Info(message string, args ...interface{}) {
	l.doLog(severityInfo, message, args...)
}

func (l *loggerImpl) Warn(message string, args ...interface{}) {
	l.doLog(severityWarn, message, args...)
}

func (l *loggerImpl) Error(message string, args ...interface{}) {
	l.doLog(severityError, message, args...)
}

func (l *loggerImpl) doLog(
	severityTag severityLevel,
	message string,
	args ...interface{},
) {
	// log to stdout/stderr first
	targetStd := os.Stdout
	if severityTag == severityError {
		targetStd = os.Stderr
	}
	_, _ = fmt.Fprintf(targetStd, "%d: [%s] %s\n", time.Now().Unix(), severityTag, fmt.Sprintf(message, args...))

	// TODO log to file, use channel to put into buffer, then have a scheduled job / threshold channel number to flush
}
