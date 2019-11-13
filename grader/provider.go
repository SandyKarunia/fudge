package grader

import (
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sandbox"
	"sync"
)

// Provider ...
func Provider(factory sandbox.Factory, logger logger.Logger) Grader {
	return &graderImpl{
		status:     StatusIdle,
		graderLock: sync.Mutex{},
		sbFactory:  factory,
		logger:     logger,
	}
}
