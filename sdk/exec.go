package sdk

import (
	"os/exec"
)

var (
	execCommand = exec.Command
)

// ExecFunctions is an interface that represents exec library in golang sdk
//go:generate mockery -name=ExecFunctions
type ExecFunctions interface {
	Command(name string, arg ...string) *exec.Cmd
}

type execFunctionsImpl struct{}

func (e *execFunctionsImpl) Command(name string, arg ...string) *exec.Cmd {
	return execCommand(name, arg...)
}

// ProvideExecFunctions ...
func ProvideExecFunctions() ExecFunctions {
	return &execFunctionsImpl{}
}
