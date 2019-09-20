package utils

import (
	"gitlab.com/sandykarunia/fudge/sdk"
)

// File ...
//go:generate mockery -name=File
type File interface {
	// Copy copies a file from source to destination
	Copy(src, dest string) error
}

type fileImpl struct {
	io sdk.IOFunctions
	os sdk.OSFunctions
}

func (f *fileImpl) Copy(src, dest string) error {
	// open the source file
	source, err := f.os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	// create destination file
	destination, err := f.os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = f.io.Copy(destination, source)
	return err
}

// ProvideFile ...
func ProvideFile(io sdk.IOFunctions, os sdk.OSFunctions) File {
	return &fileImpl{
		io: io,
		os: os,
	}
}
