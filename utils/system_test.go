package utils

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/sandykarunia/fudge/sdk/mocks"
	"testing"
)

func TestIsSudo(t *testing.T) {
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

func TestProvideSystem(t *testing.T) {
	assert.Implements(t, (*System)(nil), ProvideSystem(nil))
}
