package sdk

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestFmtFunctionsImpl_Fprintf(t *testing.T) {
	origFmtFprintf := fmtFprintf
	defer func() {
		fmtFprintf = origFmtFprintf
	}()

	called := false
	mockN := 123
	var mockErr error
	mockErr = nil
	fmtFprintf = func(
		w io.Writer,
		format string,
		a ...interface{},
	) (n int, err error) {
		called = true
		return mockN, mockErr
	}

	obj := &fmtFunctionsImpl{}

	n, err := obj.Fprintf(nil, "546", 1, "2", true)
	assert.True(t, called)
	assert.Equal(t, mockN, n)
	assert.Equal(t, mockErr, err)
}

func TestProvideFmtFunctions(t *testing.T) {
	assert.Implements(t, (*FmtFunctions)(nil), ProvideFmtFunctions())
}
