package checkers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/sandykarunia/fudge/utils/mocks"
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

func TestCheckersImpl_CheckIsolateBinaryExists(t *testing.T) {
	tests := []struct {
		desc       string
		want       bool
		fileExists bool
		fudgeDir   string
	}{
		{
			desc:       "file exists",
			want:       true,
			fileExists: true,
			fudgeDir:   "~/fudge",
		},
		{
			desc:       "file doesn't exist",
			want:       false,
			fileExists: false,
			fudgeDir:   "~/fudge",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			sysMock := &mocks.System{}
			fileMock := &mocks.File{}
			pathMock := &mocks.Path{}

			sysMock.On("GetFudgeDir").Return(test.fudgeDir)
			fileMock.On("Exists", "isolate path").Return(test.fileExists)
			pathMock.On("IsolateBinary").Return("isolate path")

			obj := &checkersImpl{sysUtils: sysMock, fileUtils: fileMock, pathUtils: pathMock}
			res := obj.CheckIsolateBinaryExists()
			assert.Equal(t, test.want, res)
		})
	}
}

func TestCheckersImpl_CheckIsolateBinaryExecutable(t *testing.T) {
	tests := []struct {
		desc       string
		want       bool
		executeErr error
	}{
		{
			desc:       "execute returns error",
			want:       false,
			executeErr: errors.New("err"),
		},
		{
			desc:       "execute does not return error",
			want:       true,
			executeErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mockPath := &mocks.Path{}
			mockPath.On("IsolateBinary").Return("isolatePath")
			mockSys := &mocks.System{}
			mockSys.On("Execute", "isolatePath", "--version").Return("", test.executeErr)

			obj := &checkersImpl{sysUtils: mockSys, pathUtils: mockPath}
			res := obj.CheckIsolateBinaryExecutable()
			assert.Equal(t, test.want, res)
		})
	}
}
