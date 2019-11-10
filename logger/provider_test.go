package logger

import (
	"github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	sysMocks := &mocks.System{}
	sysMocks.On("GetFudgeDir").Return("")
	assert.Implements(t, (*Logger)(nil), Provider(nil, sysMocks))
}
