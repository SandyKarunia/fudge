package grader

import (
	"fmt"
	"github.com/sandykarunia/fudge/grader/taskrunner"
	"github.com/sandykarunia/fudge/language"
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
	GradeAsync(payload *GradeAsyncPayload) bool
}

// GradeAsyncPayload ...
type GradeAsyncPayload struct {
	UUID               string
	SubmissionCode     string
	SubmissionLanguage language.Language
	GradingCode        string
	GradingLanguage    language.Language
	GradingMethod      Method
	MemoryLimitKB      int64
	TimeLimitMS        int64
	InputURL           []string
	OutputURL          []string
}

type graderImpl struct {
	// status of current grader
	status Status

	// a lock to prevent multiple grading at a time
	graderLock sync.Mutex

	taskRunner taskrunner.TaskRunner
	logger     logger.Logger
}

func (g *graderImpl) Status() Status {
	return g.status
}

func (g *graderImpl) GradeAsync(payload *GradeAsyncPayload) bool {
	if payload == nil {
		return false
	}

	g.logger.Info("GradeAsync triggered with UUID %s", payload.UUID)
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
	g.changeStatus(StatusAcknowledged, "Successfully triggered GradeAsync with UUID %s", payload.UUID)

	// run the grading main flow in different thread
	go g.doGrade(payload)

	return true
}

func (g *graderImpl) doGrade(payload *GradeAsyncPayload) {
	// always set to idle after everything has finished
	defer func() {
		g.changeStatus(StatusIdle, "End of doGrade function")
	}()

	// vars
	var sb sandbox.Sandbox
	var submissionCodeFilename string
	var compiledSubmissionCodeFilename string
	var inputDataFilenames []string
	var outputDataFilenames []string
	var err error

	// prepare sandbox
	g.changeStatus(StatusPrepareSandbox, "Preparing sandbox")
	if sb, err = g.taskRunner.PrepareSandbox(); err != nil {
		return
	}

	// prepare submission code
	g.changeStatus(StatusPrepareSubmissionCode, "Preparing submission code")
	if submissionCodeFilename, err = g.taskRunner.PrepareSubmissionCode(sb, payload.SubmissionCode); err != nil {
		return
	}

	// compile code first
	g.changeStatus(StatusCompiling, "Compiling source code")
	if compiledSubmissionCodeFilename, err = g.taskRunner.CompileCode(sb, submissionCodeFilename, payload.SubmissionLanguage); err != nil {
		return
	}

	g.changeStatus(StatusFetchInput, "Fetching input data")
	if inputDataFilenames, err = g.taskRunner.FetchAndWriteToFile(sb, payload.InputURL); err != nil {
		return
	}

	g.changeStatus(StatusFetchOutput, "Fetching output data")
	if outputDataFilenames, err = g.taskRunner.FetchAndWriteToFile(sb, payload.OutputURL); err != nil {
		return
	}

	// TODO grade submission (GRADING)
	g.changeStatus(StatusGrading, "Grading")

	// TODO notify result (NOTIFY_RESULT)
	g.changeStatus(StatusNotifyResult, "Notifying result via webhook HTTP request")

	g.changeStatus(StatusCleanUp, "Cleaning up sandbox with box-id = %d", sb.GetID())
	if err = g.taskRunner.CleanupSandbox(sb); err != nil {
		return
	}

	// TODO remove this, this is to prevent error during development
	fmt.Println(compiledSubmissionCodeFilename, inputDataFilenames, outputDataFilenames)
}

// a helper function to help grader change its status
func (g *graderImpl) changeStatus(nextStatus Status, message string, args ...interface{}) {
	g.logger.Info(fmt.Sprintf("%s, status: %s => %s", message, g.status, nextStatus), args...)
	g.status = nextStatus
}
