package sandbox

import (
	"fmt"
	"gitlab.com/sandykarunia/fudge/utils"
	"os"
)

const sandboxGoneErrFmt = "sandbox: id %d is gone, can't execute %s"

// Sandbox is an interface of sandbox instance, the file / folder structure inside is flat
type Sandbox interface {
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

	// GetID returns id
	GetID() int
}

type sandboxImpl struct {
	id          int
	isDestroyed bool
	path        utils.Path
}

func (s *sandboxImpl) CopyFile(source *os.File) error {
	if s.isDestroyed {
		return fmt.Errorf(sandboxGoneErrFmt, s.id, "CopyFile")
	}
	panic("implement me")
}

func (s *sandboxImpl) GetFile(fileName string) (*os.File, error) {
	if s.isDestroyed {
		return nil, fmt.Errorf(sandboxGoneErrFmt, s.id, "GetFile")
	}
	panic("implement me")
}

func (s *sandboxImpl) Run(commands ...string) error {
	if s.isDestroyed {
		return fmt.Errorf(sandboxGoneErrFmt, s.id, "Run")
	}
	panic("implement me")
}

func (s *sandboxImpl) Destroy() {
	if s.isDestroyed {
		return
	}

	s.isDestroyed = true
	panic("implement me")
}

func (s *sandboxImpl) GetID() int {
	return s.id
}
