package sandbox

import (
	"fmt"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
	"io"
	"io/ioutil"
)

const isolateCmd = "isolate"

type isolateImpl struct {
	sdkOS sdk.OSFunctions
	sdkIO sdk.IOFunctions

	id            uint32
	isDestroyed   bool
	isPrepared    bool
	isCGSupported bool
	utilsPath     utils.Path
	utilsSystem   utils.System
}

func (s *isolateImpl) WriteFile(filename string, stream io.ReadCloser) error {
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

func (s *isolateImpl) ReadFile(filename string) (string, error) {
	if !s.isActive() {
		return "", fmt.Errorf(sandboxInactiveErrFmt, s.id, "ReadFile")
	}

	// Open the file
	file, err := s.sdkOS.Open(s.utilsPath.BoxDir(s.id) + filename)
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

func (s *isolateImpl) Run(
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
		fmt.Sprintf("--fsize=%d", fileSizeLimitKB),
		fmt.Sprintf("--time=%.3f", float64(timeLimitMS)/1000.0),
	)
	if s.isCGSupported {
		isolateArgs = append(isolateArgs, "--cg", fmt.Sprintf("--cg-mem=%d", memoryLimitKB))
	} else {
		isolateArgs = append(isolateArgs, fmt.Sprintf("--mem=%d", memoryLimitKB))
	}
	if len(metaFile) > 0 {
		isolateArgs = append(isolateArgs, fmt.Sprintf("--meta=%s", metaFile))
	}
	if len(stdinFile) > 0 {
		isolateArgs = append(isolateArgs, fmt.Sprintf("--stdin=%s", stdinFile))
	}
	if len(stdoutFile) > 0 {
		isolateArgs = append(isolateArgs, fmt.Sprintf("--stdout=%s", stdoutFile))
	}
	isolateArgs = append(isolateArgs, "--", command)
	isolateArgs = append(isolateArgs, args...)

	_, err := s.utilsSystem.Execute(isolateCmd, isolateArgs...)
	return err
}

func (s *isolateImpl) Destroy() error {
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

func (s *isolateImpl) Prepare() error {
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

func (s *isolateImpl) GetID() uint32 {
	return s.id
}

// isActive returns true if the sandbox is already prepared, and not destroyed
func (s *isolateImpl) isActive() bool {
	return s.isPrepared && !s.isDestroyed
}
