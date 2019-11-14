package utils

import (
	"errors"
	sdkMocks "github.com/sandykarunia/fudge/sdk/mocks"
	utilsMocks "github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPathImpl_FudgeDir(t *testing.T) {
	tests := []struct {
		desc           string
		want           string
		userHomeDir    string
		userHomeDirErr error
	}{
		{
			desc:           "user home dir contains error, should NOT return, should continue processing",
			userHomeDir:    "this/is/home/",
			userHomeDirErr: errors.New("some error"),
			want:           "this/is/home/.fudge/",
		},
		{
			desc:        "should add '/' suffix if user home dir doesn't have it",
			userHomeDir: "this/is/home/no/slash/suffix",
			want:        "this/is/home/no/slash/suffix/.fudge/",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mockOS := &sdkMocks.OSFunctions{}
			mockSystem := &utilsMocks.System{}
			mockOS.On("UserHomeDir").Return(test.userHomeDir, test.userHomeDirErr)
			mockSystem.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return("", nil)

			obj := &pathImpl{sdkOS: mockOS, system: mockSystem}
			res := obj.FudgeDir()
			assert.Equal(t, test.want, res)
			mockSystem.AssertCalled(t, "Execute", "mkdir", "-p", test.want)
		})
	}
}

func TestProvidePath(t *testing.T) {
	assert.Implements(t, (*Path)(nil), ProvidePath(nil, nil))
}
