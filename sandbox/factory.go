package sandbox

import (
	"github.com/sandykarunia/fudge/utils"
	"math/rand"
)

// Factory is an interface for sandbox factory
type Factory interface {
	// NewSandbox produces new sandbox instance with random ID
	NewSandbox() Sandbox
}

type factoryImpl struct {
	utilsPath utils.Path
}

func (f *factoryImpl) NewSandbox() Sandbox {
	// we don't need to check if the newID exists or not, since:
	// - by right, each machine only have 1 judge
	// - the range of ID is up to 2^32-1, so if a machine have multiple judges, the chance for it to collide is small
	// - the judge will clean up the sandbox instance after usage, which means the used ID becomes available
	newID := rand.Int()
	sandbox := &sandboxImpl{
		id:        newID,
		utilsPath: f.utilsPath,
	}

	return sandbox
}
