package taskrunner

import (
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sandbox"
	"github.com/sandykarunia/fudge/utils"
	"io/ioutil"
	"strings"
)

// TaskRunner contains methods to run single task in grading process
//go:generate mockery -name=TaskRunner
type TaskRunner interface {
	// PrepareSandbox prepares the sandbox before use, returns the sandbox instance
	PrepareSandbox() (sandbox.Sandbox, error)

	// PrepareSubmissionCode writes submission code into the sandbox, returns created filename
	PrepareSubmissionCode(sb sandbox.Sandbox, submissionCode string) (filename string, err error)
}

type taskRunnerImpl struct {
	sbFactory   sandbox.Factory
	logger      logger.Logger
	utilsString utils.String
}

func (t *taskRunnerImpl) PrepareSandbox() (sandbox.Sandbox, error) {
	sb, err := t.sbFactory.NewPreparedSandbox()
	if err != nil {
		t.logger.Error("Failed to prepare sandbox, err = %s", err.Error())
		return nil, err
	}
	t.logger.Info("Sandbox prepared with box-id = %d", sb.GetID())
	return sb, nil
}

func (t *taskRunnerImpl) PrepareSubmissionCode(sb sandbox.Sandbox, submissionCode string) (filename string, err error) {
	submissionCodeFilename := t.utilsString.GenerateRandomAlphanumeric(16)
	if err = sb.WriteFile(submissionCodeFilename, ioutil.NopCloser(strings.NewReader(submissionCode))); err != nil {
		t.logger.Error("Failed to prepare submission code, err = %s", err.Error())
	}
	t.logger.Info("Submission code prepared with filename = %s", submissionCodeFilename)
	return submissionCodeFilename, err
}
