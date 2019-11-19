package grader

import (
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sandbox"
	"github.com/sandykarunia/fudge/utils"
	"sync"
)

// Provider ...
func Provider(factory sandbox.Factory, logger logger.Logger, utilsString utils.String) Grader {
	return &graderImpl{
		status:      StatusIdle,
		graderLock:  sync.Mutex{},
		sbFactory:   factory,
		logger:      logger,
		utilsString: utilsString,
	}
}
