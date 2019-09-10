package sandbox

import (
	"math/rand"
	"os"
)

// Sandbox is an interface of sandbox, the file / folder structure inside is flat
type Sandbox interface {
	// Initialize prepares the sandbox before it is used
	Initialize() error

	// CopyFile copies the source file into the sandbox
	// we don't need the target path in sandbox because the file / folder structure is flat
	CopyFile(source *os.File) error

	// GetFile gets a file from
	// we only need the filename because the file / folder structure is flat
	GetFile(fileName string) (*os.File, error)

	// Run runs commands inside the sandbox instance
	Run(commands ...string) error

	// Destroy the sandbox instance, after it is destroyed, we should not use the sandbox anymore
	Destroy()
}

// New creates a new sandbox instance
func New() Sandbox {
	return &moeSandbox{
		isDestroyed: false,
		boxID:       rand.Intn(1000), // it is ok if we do this here, by right one machine only run one judge
	}
}
