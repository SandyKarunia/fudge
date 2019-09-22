package utils

import (
	"gitlab.com/sandykarunia/fudge/sdk"
)

// File ...
//go:generate mockery -name=File
type File interface {
	// Copy copies a file from source to destination
	Copy(src, dest string) error

	// Exists checks if the path exists AND it is a file
	Exists(path string) bool
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

func (f *fileImpl) Exists(path string) bool {
	info, err := f.os.Stat(path)

	if f.os.IsNotExist(err) {
		return false
	}
	if info.IsDir() {
		return false
	}

	return true
}

// ProvideFile ...
func ProvideFile(io sdk.IOFunctions, os sdk.OSFunctions) File {
	return &fileImpl{
		io: io,
		os: os,
	}
}
