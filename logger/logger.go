package logger

import (
	"fmt"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
	"os"
	"sync"
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

	// FlushBuffer flushes the message buffer into a log file
	FlushBuffer()
}

type loggerImpl struct {
	os  sdk.OSFunctions
	sys utils.System

	_logFileLock   sync.Mutex
	_logFilePath   string
	_logFileBuffer chan string
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

func (l *loggerImpl) FlushBuffer() {
	// lock the file when we want to use it as there could be multiple threads that call this function
	l._logFileLock.Lock()
	defer l._logFileLock.Unlock()

	var allLogsInBuffer []string
	// get all logs inside the buffer
	// this for loop should never call any logging action, otherwise it will cause infinite loop
	for {
		var done bool
		select {
		case log := <-l._logFileBuffer:
			allLogsInBuffer = append(allLogsInBuffer, log)
		default:
			done = true
			break
		}

		if done {
			break
		}
	}

	// if no logs, then just return
	if len(allLogsInBuffer) == 0 {
		return
	}

	// write to log file
	logFile, err := os.OpenFile(l._logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		l.Warn("Failed to open log file, err = %s\n", err.Error())
		return
	}
	defer logFile.Close()
	for _, log := range allLogsInBuffer {
		_, err = logFile.WriteString(log)
		if err != nil {
			l.Warn("Failed to append log to file, err = %s\n", err.Error())
		}
	}
}

func (l *loggerImpl) doLog(
	severityTag severityLevel,
	message string,
	args ...interface{},
) {
	currentTime := time.Now()
	logMsg := fmt.Sprintf(
		"%s [%s] %s\n", currentTime.Format(time.RFC3339), severityTag, fmt.Sprintf(message, args...),
	)

	// log to stdout/stderr first
	targetStd := os.Stdout
	if severityTag == severityError {
		targetStd = os.Stderr
	}
	_, _ = fmt.Fprintf(targetStd, logMsg)

	// put into buffer for log file
	l._logFileBuffer <- logMsg
}

func (l *loggerImpl) flushBufferJob() {
	// trigger flush every 1 minute
	for {
		<-time.After(1 * time.Minute)
		l.FlushBuffer()
	}
}

func (l *loggerImpl) init() Logger {
	l._logFilePath = fmt.Sprintf("%slog-%d.txt", l.sys.GetFudgeDir(), time.Now().Unix())
	l._logFileBuffer = make(chan string, 1000)
	go l.flushBufferJob()

	return l
}
