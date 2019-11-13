package grader

import (
	"fmt"
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sandbox"
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
		uuid string,
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

	// factory for sandbox
	sbFactory sandbox.Factory

	// deps
	logger logger.Logger
}

func (g *graderImpl) Status() Status {
	return g.status
}

func (g *graderImpl) GradeAsync(
	uuid string,
	submissionCode, gradingCode string,
	gradingMethod Method,
	memoryLimitKB, timeLimitMS int64,
	inputURL, outputURL []string) bool {
	g.logger.Info("GradeAsync triggered with uuid %s", uuid)
	// check grader status first, if there is another grader that is running (i.e. status != idle), then return false
	// use check-lock-check pattern
	if g.status != StatusIdle {
		return false
	}
	g.graderLock.Lock()
	defer g.graderLock.Unlock()
	if g.status != StatusIdle {
		return false
	}

	// change grader status first before return
	g.changeStatus(StatusAcknowledged, "Successfully triggered GradeAsync with uuid %s", uuid)

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
		g.changeStatus(StatusIdle, "End of doGrade function")
	}()

	// all operations below are inside the sandbox

	g.changeStatus(StatusPrepare, "Preparing sandbox")
	sb := g.sbFactory.NewPreparedSandbox()
	g.logger.Info("Sandbox prepared with box-id = %d", sb.GetID())

	// TODO fetch input, put into file (FETCH_INPUT)
	g.changeStatus(StatusFetchInput, "Fetching input data")

	// TODO fetch output, put into file (FETCH_OUTPUT)
	g.changeStatus(StatusFetchOutput, "Fetching output data")

	// TODO grade submission (GRADING)
	g.changeStatus(StatusGrading, "Grading")

	// TODO notify result (NOTIFY_RESULT)
	g.changeStatus(StatusNotifyResult, "Notifying result via webhook HTTP request")

	// TODO cleanup sandbox (CLEAN_UP)
	//g.changeStatus(StatusCleanUp, "Cleaning up sandbox with box-id = %d", sb.GetID())

}

// a helper function to help grader change its status
func (g *graderImpl) changeStatus(nextStatus Status, message string, args ...interface{}) {
	g.logger.Info(fmt.Sprintf("%s, status: %s => %s", message, g.status, nextStatus), args...)
	g.status = nextStatus
}
