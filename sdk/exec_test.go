package sdk

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestExecFunctionsImpl_Command(t *testing.T) {
	// mock execCommand
	originalExecCommand := execCommand
	execCommand = func(name string, arg ...string) *exec.Cmd {
		return nil
	}
	defer func() {
		execCommand = originalExecCommand
	}()

	obj := &execFunctionsImpl{}
	res := obj.Command("")
	assert.Nil(t, res)
}

func TestProvideExecFunctions(t *testing.T) {
	assert.Implements(t, (*ExecFunctions)(nil), ProvideExecFunctions())
}
