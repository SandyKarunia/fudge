package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvideStdLogger(t *testing.T) {
	assert.Implements(t, (*Logger)(nil), ProvideStdLogger(nil))
}
