package sandbox

import (
	"io"
)

const sandboxInactiveErrFmt = "sandbox: id %d is inactive, can't execute %s"

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
	//   it can be empty which means no stdin file
	// - stdoutFile = redirect stdout to a file, it has to be accessible from inside the sandbox, by default we redirect
	//   stderr to stdout. It can be empty which means no stdout file
	// - metaFile = output metadata (extra information about the run) to a file.
	//   It can be empty, which means no meta file
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
