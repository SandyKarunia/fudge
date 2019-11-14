package utils

import (
	"errors"
	sdkMocks "github.com/sandykarunia/fudge/sdk/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathImpl_IsolateBinary(t *testing.T) {
	obj := &pathImpl{}
	assert.Equal(t, "/usr/local/bin/isolate", obj.IsolateBinary())
}

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
			mockOS.On("UserHomeDir").Return(test.userHomeDir, test.userHomeDirErr)

			obj := &pathImpl{sdkOS: mockOS}
			res := obj.FudgeDir()
			assert.Equal(t, test.want, res)
		})
	}
}

func TestProvidePath(t *testing.T) {
	assert.Implements(t, (*Path)(nil), ProvidePath(nil, nil))
}
