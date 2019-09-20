package sdk

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestIoFunctionsImpl_Copy(t *testing.T) {
	// mock ioCopy
	originalIOCopy := ioCopy
	ioCopy = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return -1, nil
	}
	defer func() {
		ioCopy = originalIOCopy
	}()

	obj := &ioFunctionsImpl{}
	res, reserr := obj.Copy(nil, nil)
	assert.Nil(t, reserr)
	assert.Equal(t, int64(-1), res)
}

func TestProvideIOFunctions(t *testing.T) {
	assert.Implements(t, (*IOFunctions)(nil), ProvideIOFunctions())
}
