package utils

import (
	"github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathImpl_IsolateBinary(t *testing.T) {
	mockSystem := &mocks.System{}
	mockSystem.On("GetFudgeDir").Return("dir/")

	obj := &pathImpl{system: mockSystem}
	assert.Equal(t, "dir/isolate", obj.IsolateBinary())
}

func TestProvidePath(t *testing.T) {
	assert.Implements(t, (*Path)(nil), ProvidePath(nil))
}
