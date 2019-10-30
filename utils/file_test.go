package utils

import (
	"errors"
	"github.com/sandykarunia/fudge/sdk/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func TestFileImpl_Copy(t *testing.T) {
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

func TestFileImpl_Exists(t *testing.T) {
	tests := []struct {
		desc          string
		fileInfoIsDir bool
		isNotExist    bool
		want          bool
	}{
		{
			desc:       "file does not exist",
			want:       false,
			isNotExist: true,
		},
		{
			desc:          "file exists, but is a directory",
			want:          false,
			isNotExist:    false,
			fileInfoIsDir: true,
		},
		{
			desc:          "file exists, and not a directory",
			want:          true,
			isNotExist:    false,
			fileInfoIsDir: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			osMock := &mocks.OSFunctions{}
			fileInfoMock := &mocks.FileInfo{}

			osMock.On("Stat", "the-path").Return(fileInfoMock, nil)
			osMock.On("IsNotExist", mock.Anything).Return(test.isNotExist)
			fileInfoMock.On("IsDir").Return(test.fileInfoIsDir)

			obj := &fileImpl{os: osMock}
			res := obj.Exists("the-path")
			assert.Equal(t, test.want, res)
		})
	}
}

func TestProvideFile(t *testing.T) {
	res := ProvideFile(nil, nil)
	assert.Implements(t, (*File)(nil), res)
}
