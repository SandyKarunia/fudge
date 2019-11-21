package sandbox

import (
	"fmt"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
	"io"
	"io/ioutil"
	"os"
)

const sandboxInactiveErrFmt = "sandbox: id %d is inactive, can't execute %s"
const isolateCmd = "isolate"

// Sandbox is an interface of sandbox instance, the file / folder structure inside is flat
type Sandbox interface {
	// WriteFile writes a file into the sandbox with specified filename
	// if the file exists, it will overwrite the file
	WriteFile(filename string, stream io.ReadCloser) error

	// ReadFile reads a file content and returns it as string
	ReadFile(fileName string) (string, error)

	// Run runs a command inside the sandbox instance
	// - timeLimitMS = the run-time limit for the program to run in milliseconds,
	//   also limits the wall-time to be 2*timeLimitMs
	// - memoryLimitKB = limit memory usage of the program in KB
	// - fileSizeLimitKB = limit file size created by the program, it supports multiple files
	// - stdinFile = redirect stdin from file, it has to be accessible from inside the sandbox,
	//   it can be empty which means no stdin
	// - stdoutFile = redirect stdout to a file, it has to be accessible from inside the sandbox, by default we redirect
	//   stderr to stdout.
	// - metaFile = output metadata (extra information about the run) to a file
	Run(
		timeLimitMS, memoryLimitKB, fileSizeLimitKB int64,
		stdinFile, stdoutFile, metaFile string,
		command string, args ...string,
	) error

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

func (s *sandboxImpl) ReadFile(fileName string) (string, error) {
	if !s.isActive() {
		return "", fmt.Errorf(sandboxInactiveErrFmt, s.id, "ReadFile")
	}

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read all bytes, it is safe because we already put limit on file size when we run the program, and normally the
	// limit is not that big
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (s *sandboxImpl) Run(
	timeLimitMS, memoryLimitKB, fileSizeLimitKB int64,
	stdinFile, stdoutFile, metaFile string,
	command string, args ...string,
) error {
	if !s.isActive() {
		return fmt.Errorf(sandboxInactiveErrFmt, s.id, "Run")
	}

	var isolateArgs []string
	isolateArgs = append(isolateArgs,
		fmt.Sprintf("--box-id=%d", s.id),
		fmt.Sprintf("--meta=%s", metaFile),
		fmt.Sprintf("--stdout=%s", stdoutFile),
		fmt.Sprintf("--fsize=%d", fileSizeLimitKB),
		fmt.Sprintf("--time=%.3f", float64(timeLimitMS)/1000.0),
	)
	if s.isCGSupported {
		isolateArgs = append(isolateArgs, "--cg", fmt.Sprintf("--cg-mem=%d", memoryLimitKB))
	} else {
		isolateArgs = append(isolateArgs, fmt.Sprintf("--mem=%d", memoryLimitKB))
	}
	if len(stdinFile) > 0 {
		isolateArgs = append(isolateArgs, fmt.Sprintf("--stdin=%s", stdinFile))
	}
	isolateArgs = append(isolateArgs, "--", command)
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
