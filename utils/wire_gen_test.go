package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileInstance(t *testing.T) {
	obj := FileInstance()
	assert.Implements(t, (*File)(nil), obj)
}

func TestSignatureInstance(t *testing.T) {
	obj := SignatureInstance()
	assert.Implements(t, (*Signature)(nil), obj)
}

func TestStringInstance(t *testing.T) {
	obj := StringInstance()
	assert.Implements(t, (*String)(nil), obj)
}

func TestSystemInstance(t *testing.T) {
	obj := SystemInstance()
	assert.Implements(t, (*System)(nil), obj)
}
