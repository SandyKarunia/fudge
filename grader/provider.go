package grader

import (
	"github.com/sandykarunia/fudge/grader/taskrunner"
	"github.com/sandykarunia/fudge/logger"
	"sync"
)

// Provider ...
func Provider(logger logger.Logger, tr taskrunner.TaskRunner) Grader {
	return &graderImpl{
		status:     StatusIdle,
		graderLock: sync.Mutex{},
		logger:     logger,
		taskRunner: tr,
	}
}
