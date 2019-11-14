package utils

import (
	"bytes"
	"github.com/sandykarunia/fudge/sdk/mocks"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestSystemImpl_IsSudo(t *testing.T) {
	tests := []struct {
		euid        int
		sudoUIDEnv  string
		sudoGIDEnv  string
		sudoUserEnv string
		want        bool
		desc        string
	}{
		{
			euid: 123,
			want: false,
			desc: "euid is not 0",
		},
		{
			sudoUIDEnv:  "a",
			sudoGIDEnv:  "b",
			sudoUserEnv: "c",
			want:        true,
			desc:        "euid is 0 and all env variables are available",
		},
		{
			sudoUIDEnv:  "a",
			sudoUserEnv: "c",
			want:        false,
			desc:        "euid is 0 but gid env is not available",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mockOS := &mocks.OSFunctions{}

			mockOS.On("Geteuid").Return(test.euid)
			mockOS.On("Getenv", "SUDO_UID").Return(test.sudoUIDEnv)
			mockOS.On("Getenv", "SUDO_GID").Return(test.sudoGIDEnv)
			mockOS.On("Getenv", "SUDO_USER").Return(test.sudoUserEnv)

			obj := &systemImpl{os: mockOS}
			res := obj.IsSudo()
			assert.Equal(t, test.want, res)
		})
	}
}

func TestSystemImpl_VerifyPkgInstalled(t *testing.T) {
	tests := []struct {
		wantError bool
		desc      string
		cmd       *exec.Cmd
	}{
		{
			desc:      "deliberately set stderr in exec.Cmd, so CombinedOutput will return error",
			wantError: true,
			cmd: &exec.Cmd{
				Stderr: &bytes.Buffer{},
			},
		},
		{
			desc:      "empty, normal exec.Cmd",
			wantError: false,
			cmd:       exec.Command("ls"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mockExec := &mocks.ExecFunctions{}
			mockExec.On("Command", "dpkg", "-s", "pkg").Return(test.cmd)
			obj := &systemImpl{exec: mockExec}
			res := obj.VerifyPkgInstalled("pkg")
			if test.wantError {
				assert.Error(t, res)
			} else {
				assert.NoError(t, res)
			}
		})
	}
}

func TestProvideSystem(t *testing.T) {
	assert.Implements(t, (*System)(nil), ProvideSystem(nil, nil))
}
