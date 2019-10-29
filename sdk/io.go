package sdk

import (
	"io"
)

var (
	ioCopy = io.Copy
)

// IOFunctions is an interface that represents io library in golang sdk
//go:generate mockery -name=IOFunctions
type IOFunctions interface {
	Copy(dest io.Writer, src io.Reader) (written int64, err error)
}

type ioFunctionsImpl struct{}

func (i *ioFunctionsImpl) Copy(dest io.Writer, src io.Reader) (written int64, err error) {
	return ioCopy(dest, src)
}

// ProvideIOFunctions ...
func ProvideIOFunctions() IOFunctions {
	return &ioFunctionsImpl{}
}
