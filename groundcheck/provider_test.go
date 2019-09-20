package groundcheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	assert.Implements(t, (*GroundCheck)(nil), Provider(nil))
}
