package sandbox

import (
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
}

func (f *factoryImpl) NewPreparedSandbox() Sandbox {
	// we don't need to check if the newID exists or not, since:
	// - by right, each machine only have 1 judge
	// - the range of ID is 0-999
	// - the judge will clean up the sandbox instance after usage, which means the used ID becomes available
	newID := rand.Uint32() % 1000
	sandbox := &sandboxImpl{
		sdkOS:       f.sdkOS,
		sdkIO:       f.sdkIO,
		id:          newID,
		isDestroyed: false,
		isPrepared:  false,
		utilsPath:   f.utilsPath,
		utilsSystem: f.utilsSystem,
	}
	sandbox.Prepare()

	return sandbox
}
