package grader

import (
	"sync"
)

// Grader is the main component of this judge that will handle a submission at a time
//go:generate mockery -name=Grader
type Grader interface {
	// Status returns the current status of the grader
	Status() Status

	// GradeAsync is the main body of the grader which will grade the requested submission asynchronously.
	// Checks the status of the grader first before grading, returns false if it is not idle
	GradeAsync(
		submissionCode, gradingCode string,
		gradingMethod Method,
		memoryLimitKB, timeLimitMS int64,
		inputURL, outputURL []string) bool
}

type graderImpl struct {
	// status of current grader
	status Status

	// a lock to prevent multiple grading at a time
	graderLock sync.Mutex
}

func (g *graderImpl) Status() Status {
	return g.status
}

func (g *graderImpl) GradeAsync(
	submissionCode, gradingCode string,
	gradingMethod Method,
	memoryLimitKB, timeLimitMS int64,
	inputURL, outputURL []string) bool {
	// check judge status first, if there is another grader that is running (i.e. status != idle), then return false
	// use check-lock-check pattern
	if g.status != StatusIdle {
		return false
	}
	g.graderLock.Lock()
	defer g.graderLock.Unlock()
	if g.status != StatusIdle {
		return false
	}

	// change judge status first before return
	g.status = StatusAcknowledged

	// run the grading main flow in different thread
	go g.doGrade(submissionCode, gradingCode, gradingMethod, memoryLimitKB, timeLimitMS, inputURL, outputURL)

	return true
}

func (g *graderImpl) doGrade(
	submissionCode, gradingCode string,
	gradingMethod Method,
	memoryLimitKB, timeLimitMS int64,
	inputURL, outputURL []string) {
	// always set to idle after everything has finished
	defer func() {
		g.status = StatusIdle
	}()

	// all operations below are inside the sandbox

	// TODO prepare sandbox (PREPARE)
	g.status = StatusPrepare
	// TODO fetch input, put into file (FETCH_INPUT)
	g.status = StatusFetchInput
	// TODO fetch output, put into file (FETCH_OUTPUT)
	g.status = StatusFetchOutput
	// TODO grade submission (GRADING)
	g.status = StatusGrading
	// TODO notify result (NOTIFY_RESULT)
	g.status = StatusNotifyResult
	// TODO cleanup sandbox (CLEAN_UP)
	g.status = StatusCleanUp
}
