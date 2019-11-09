package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	assert.Implements(t, (*Logger)(nil), Provider(nil))
}
