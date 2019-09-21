package checkers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	assert.Implements(t, (*Checkers)(nil), Provider(nil, nil))
}
