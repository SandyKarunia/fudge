package sandbox

import (
	"fmt"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
	"io"
	"io/ioutil"
)

type fakeImpl struct {
	sdkOS   sdk.OSFunctions
	sdkIO   sdk.IOFunctions
	sdkExec sdk.ExecFunctions

	id          uint32
	isDestroyed bool
	isPrepared  bool
	utilsPath   utils.Path
	utilsSystem utils.System
}

func (f *fakeImpl) WriteFile(filename string, stream io.ReadCloser) error {
	if !f.isActive() {
		return fmt.Errorf(sandboxInactiveErrFmt, f.id, "WriteFile")
	}

	// Create / Open the file
	out, err := f.sdkOS.Create(f.getFakeBoxPath() + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write stream into the file
	_, err = f.sdkIO.Copy(out, stream)
	return err
}

func (f *fakeImpl) ReadFile(fileName string) (string, error) {
	if !f.isActive() {
		return "", fmt.Errorf(sandboxInactiveErrFmt, f.id, "ReadFile")
	}

	// Open the file
	file, err := f.sdkOS.Open(f.getFakeBoxPath() + fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// read all bytes
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (f *fakeImpl) Run(
	timeLimitMS, memoryLimitKB, fileSizeLimitKB int64,
	stdinFile, stdoutFile, metaFile string,
	command string, args ...string,
) error {
	if !f.isActive() {
		return fmt.Errorf(sandboxInactiveErrFmt, f.id, "Run")
	}

	// ignore all limits and proceed to run the command directly
	// prepare command
	cmd := f.sdkExec.Command(command, args...)

	// change command directory to box dir
	cmd.Dir = f.getFakeBoxPath()

	// in file
	inString, _ := f.ReadFile(stdinFile)
	stdinPipe, _ := cmd.StdinPipe()
	io.WriteString(stdinPipe, inString)

	// out file
	out, _ := f.sdkOS.Create(f.getFakeBoxPath() + stdoutFile)
	defer out.Close()
	cmd.Stdout = out
	cmd.Stderr = out

	// create empty meta file
	meta, _ := f.sdkOS.Create(f.getFakeBoxPath() + metaFile)
	defer meta.Close()

	// run the cmd
	return cmd.Run()
}

func (f *fakeImpl) Destroy() error {
	if !f.isActive() || f.isDestroyed {
		return nil
	}
	f.isDestroyed = true

	// destroy / cleanup the fake sandbox by simply removing the folder
	_, err := f.utilsSystem.Execute("rm", "-r", f.getFakeBoxPath())
	return err
}

func (f *fakeImpl) Prepare() error {
	if f.isPrepared || f.isDestroyed {
		return nil
	}
	f.isPrepared = true

	// create fake sandbox by simply creating the folder
	_, err := f.utilsSystem.Execute("mkdir", "-p", f.getFakeBoxPath())
	return err
}

func (f *fakeImpl) GetID() uint32 {
	return f.id
}

// isActive returns true if the sandbox is already prepared, and not destroyed
func (f *fakeImpl) isActive() bool {
	return f.isPrepared && !f.isDestroyed
}

func (f *fakeImpl) getFakeBoxPath() string {
	return fmt.Sprintf("%sfake_sandbox/%d/box/", f.utilsPath.FudgeDir(), f.id)
}
