package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIOInstance(t *testing.T) {
	obj := IOInstance()
	assert.Implements(t, (*IOFunctions)(nil), obj)
}

func TestOSInstance(t *testing.T) {
	obj := OSInstance()
	assert.Implements(t, (*OSFunctions)(nil), obj)
}
