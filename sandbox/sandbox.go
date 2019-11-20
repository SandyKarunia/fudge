package sandbox

import (
	"fmt"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
	"io"
	"os"
)

const sandboxInactiveErrFmt = "sandbox: id %d is inactive, can't execute %s"
const isolateCmd = "isolate"

// Sandbox is an interface of sandbox instance, the file / folder structure inside is flat
type Sandbox interface {
	// WriteFile writes a file into the sandbox with specified filename
	// if the file exists, it will overwrite the file
	WriteFile(filename string, stream io.ReadCloser) error

	// GetFile gets a file from
	// we only need the filename because the file / folder structure is flat
	GetFile(fileName string) (*os.File, error)

	// Run runs commands inside the sandbox instance
	Run(commands string, args ...string) error

	// Destroy the sandbox instance, after it is destroyed, we should not use the sandbox anymore
	Destroy() error

	// Prepare the sandbox instance, it has to be prepared first before the sandbox is used
	Prepare() error

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
}

func (s *sandboxImpl) WriteFile(filename string, stream io.ReadCloser) error {
	if !s.isActive() {
		return fmt.Errorf(sandboxInactiveErrFmt, s.id, "WriteFile")
	}

	// Create / Open the file
	out, err := s.sdkOS.Create(s.utilsPath.BoxDir(s.id) + filename)
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

func (s *sandboxImpl) Run(command string, args ...string) error {
	if !s.isActive() {
		return fmt.Errorf(sandboxInactiveErrFmt, s.id, "Run")
	}

	var isolateArgs []string
	isolateArgs = append(isolateArgs,
		fmt.Sprintf("--box-id=%d", s.id),
		"--run",
		"--",
		command,
	)
	isolateArgs = append(isolateArgs, args...)

	_, err := s.utilsSystem.Execute(isolateCmd, isolateArgs...)
	return err
}

func (s *sandboxImpl) Destroy() error {
	if !s.isActive() || s.isDestroyed {
		return nil
	}
	s.isDestroyed = true

	// destroy / cleanup the sandbox
	var args []string
	args = append(args,
		fmt.Sprintf("--box-id=%d", s.id),
		"--cleanup",
	)
	_, err := s.utilsSystem.Execute(isolateCmd, args...)
	return err
}

func (s *sandboxImpl) Prepare() error {
	if s.isPrepared || s.isDestroyed {
		return nil
	}
	s.isPrepared = true

	// create sandbox
	var args []string
	if s.isCGSupported {
		args = append(args, "--cg")
	}
	args = append(args,
		fmt.Sprintf("--box-id=%d", s.id),
		"--init",
	)
	_, err := s.utilsSystem.Execute(isolateCmd, args...)
	return err
}

func (s *sandboxImpl) GetID() uint32 {
	return s.id
}

// isActive returns true if the sandbox is already prepared, and not destroyed
func (s *sandboxImpl) isActive() bool {
	return s.isPrepared && !s.isDestroyed
}
