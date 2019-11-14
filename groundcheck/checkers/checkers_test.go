package checkers

import (
	"errors"
	"github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCheckersImpl_CheckSudo(t *testing.T) {
	tests := []struct {
		desc   string
		isSudo bool
		want   bool
	}{
		{
			desc:   "current environment is not sudo",
			isSudo: false,
			want:   false,
		},
		{
			desc:   "current environment is sudo",
			isSudo: true,
			want:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			sysMock := &mocks.System{}
			sysMock.On("IsSudo").Return(test.isSudo)

			obj := &checkersImpl{sysUtils: sysMock}
			res := obj.CheckSudo()
			assert.Equal(t, test.want, res)
		})
	}
}

func TestCheckersImpl_CheckLibcapDevPkg(t *testing.T) {
	tests := []struct {
		desc         string
		verifyPkgErr error
		want         bool
	}{
		{
			desc:         "verify package returns error",
			verifyPkgErr: errors.New("err"),
			want:         false,
		},
		{
			desc:         "verify package returns nil",
			verifyPkgErr: nil,
			want:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			sysMock := &mocks.System{}
			sysMock.On("VerifyPkgInstalled", "libcap-dev").Return(test.verifyPkgErr)

			obj := &checkersImpl{sysUtils: sysMock}
			res := obj.CheckLibcapDevPkg()
			assert.Equal(t, test.want, res)
		})
	}
}

func TestCheckersImpl_CheckIsolateBinaryValid(t *testing.T) {
	tests := []struct {
		desc       string
		want       bool
		executeErr error
	}{
		{
			desc:       "binary valid",
			want:       true,
			executeErr: nil,
		},
		{
			desc:       "binary invalid",
			want:       false,
			executeErr: errors.New("err"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			sysMock := &mocks.System{}

			sysMock.On("Execute", mock.Anything, mock.Anything).Return("", test.executeErr)

			obj := &checkersImpl{sysUtils: sysMock}
			res := obj.CheckIsolateBinaryValid()
			assert.Equal(t, test.want, res)
		})
	}
}
