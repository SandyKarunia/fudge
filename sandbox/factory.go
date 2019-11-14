package sandbox

import (
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
	"math/rand"
)

// Factory is an interface for sandbox factory
type Factory interface {
	// NewPreparedSandbox produces new prepared sandbox instance with random ID
	NewPreparedSandbox() Sandbox
}

type factoryImpl struct {
	sdkOS       sdk.OSFunctions
	sdkIO       sdk.IOFunctions
	utilsPath   utils.Path
	utilsSystem utils.System
	logger      logger.Logger

	isCGSupported bool
}

func (f *factoryImpl) NewPreparedSandbox() Sandbox {
	// we don't need to check if the newID exists or not, since:
	// - by right, each machine only have 1 judge
	// - the range of ID is 0-999
	// - the judge will clean up the sandbox instance after usage, which means the used ID becomes available
	newID := rand.Uint32() % 1000
	sandbox := &sandboxImpl{
		sdkOS:         f.sdkOS,
		sdkIO:         f.sdkIO,
		id:            newID,
		isDestroyed:   false,
		isPrepared:    false,
		isCGSupported: f.isCGSupported,
		utilsPath:     f.utilsPath,
		utilsSystem:   f.utilsSystem,
	}
	sandbox.Prepare()

	return sandbox
}

func (f *factoryImpl) init() {
	// check if control group is supported or not
	if len(f.sdkOS.Getenv("CONFIG_CPUSETS")) > 0 {
		f.isCGSupported = true
		f.logger.Info("CG is supported in current machine")
	} else {
		f.isCGSupported = false
		f.logger.Warn("CG is not supported in current machine")
	}
}
