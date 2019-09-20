package utils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/sandykarunia/fudge/sdk/mocks"
	"os"
	"testing"
)

func TestFileUtilsImpl_Copy(t *testing.T) {
	tests := []struct {
		openErr     error
		createErr   error
		copyErr     error
		source      *os.File
		dest        *os.File
		wantReturn  error
		description string
	}{
		{
			openErr:     errors.New("open error"),
			wantReturn:  errors.New("open error"),
			description: "Open file error",
		},
		{
			source:      &os.File{},
			createErr:   errors.New("create error"),
			wantReturn:  errors.New("create error"),
			description: "Create file error",
		},
		{
			source:      &os.File{},
			dest:        &os.File{},
			copyErr:     errors.New("copy error"),
			wantReturn:  errors.New("copy error"),
			description: "Copy file error",
		},
		{
			source:      &os.File{},
			dest:        &os.File{},
			description: "No error",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockOS := &mocks.OSFunctions{}
			mockIO := &mocks.IOFunctions{}

			paramSrc := "a"
			paramDest := "b"

			mockOS.On("Open", paramSrc).Return(test.source, test.openErr)
			mockOS.On("Create", paramDest).Return(test.dest, test.createErr)

			mockIO.On("Copy", mock.Anything, mock.Anything).Return(int64(0), test.copyErr)

			obj := &fileImpl{
				os: mockOS,
				io: mockIO,
			}
			ret := obj.Copy(paramSrc, paramDest)

			assert.Equal(t, test.wantReturn, ret)
		})
	}
}

func TestProvideFile(t *testing.T) {
	res := ProvideFile(nil, nil)
	assert.Implements(t, (*File)(nil), res)
}
