package sandbox

import (
	"fmt"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
	"io"
	"os"
)

const sandboxInactiveErrFmt = "sandbox: id %d is inactive, can't execute %s"

// Sandbox is an interface of sandbox instance, the file / folder structure inside is flat
type Sandbox interface {
	// WriteFile writes a file into the sandbox with specified filename
	// if the file exists, it will overwrite the file
	WriteFile(filename string, stream io.ReadCloser) error

	// GetFile gets a file from
	// we only need the filename because the file / folder structure is flat
	GetFile(fileName string) (*os.File, error)

	// Run runs commands inside the sandbox instance
	Run(commands ...string) error

	// Destroy the sandbox instance, after it is destroyed, we should not use the sandbox anymore
	Destroy()

	// Prepare the sandbox instance, it has to be prepared first before the sandbox is used
	Prepare()

	// GetID returns id
	GetID() uint32
}

type sandboxImpl struct {
	sdkOS sdk.OSFunctions
	sdkIO sdk.IOFunctions

	id            uint32
	isDestroyed   bool
	isPrepared    bool
	isCGSupported bool
	utilsPath     utils.Path
	utilsSystem   utils.System
	sandboxDir    string
}

func (s *sandboxImpl) WriteFile(filename string, stream io.ReadCloser) error {
	if !s.isActive() {
		return fmt.Errorf(sandboxInactiveErrFmt, s.id, "WriteFile")
	}

	// Create / Open the file
	out, err := s.sdkOS.Create(s.sandboxDir + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write stream into the file
	_, err = s.sdkIO.Copy(out, stream)
	return err
}

func (s *sandboxImpl) GetFile(fileName string) (*os.File, error) {
	if !s.isActive() {
		return nil, fmt.Errorf(sandboxInactiveErrFmt, s.id, "GetFile")
	}
	panic("implement me")
}

func (s *sandboxImpl) Run(commands ...string) error {
	if !s.isActive() {
		return fmt.Errorf(sandboxInactiveErrFmt, s.id, "Run")
	}
	panic("implement me")
}

func (s *sandboxImpl) Destroy() {
	if !s.isActive() || s.isDestroyed {
		return
	}
	s.isDestroyed = true

	// destroy / cleanup the sandbox
	out, err := s.utilsSystem.Execute(
		"isolate", fmt.Sprintf("--box-id=%d", s.id), "--cleanup",
	)
	// TODO dont print like this
	fmt.Println(out)
	fmt.Println(err)

	panic("implement me")
}

func (s *sandboxImpl) Prepare() {
	if s.isPrepared || s.isDestroyed {
		return
	}
	s.isPrepared = true

	// create sandbox
	out, err := s.utilsSystem.Execute(
		"isolate", s.cgOption(), fmt.Sprintf("--box-id=%d", s.id), "--init",
	)
	// TODO dont print like this, put sandbox directory to sandboxDir variable
	fmt.Println(out)
	fmt.Println(err)
}

func (s *sandboxImpl) GetID() uint32 {
	return s.id
}

// isActive returns true if the sandbox is already prepared, and not destroyed
func (s *sandboxImpl) isActive() bool {
	return s.isPrepared && !s.isDestroyed
}

func (s *sandboxImpl) cgOption() string {
	if s.isCGSupported {
		return "--cg"
	}
	return ""
}
