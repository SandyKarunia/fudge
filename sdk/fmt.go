package sdk

import (
	"fmt"
	"io"
)

var (
	fmtFprintf = fmt.Fprintf
)

// FmtFunctions is an interface that represents fmt library in golang sdk
//go:generate mockery -name=FmtFunctions
type FmtFunctions interface {
	Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
}

type fmtFunctionsImpl struct{}

func (f *fmtFunctionsImpl) Fprintf(
	w io.Writer,
	format string,
	a ...interface{},
) (n int, err error) {
	return fmtFprintf(w, format, a)
}

// ProvideFmtFunctions ...
func ProvideFmtFunctions() FmtFunctions {
	return &fmtFunctionsImpl{}
}
