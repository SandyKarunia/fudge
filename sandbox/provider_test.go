package sandbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	assert.Implements(t, (*Factory)(nil), Provider(nil))
}
