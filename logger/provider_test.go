package logger

import (
	"github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProvider(t *testing.T) {
	pathMocks := &mocks.Path{}
	pathMocks.On("FudgeDir").Return("")
	assert.Implements(t, (*Logger)(nil), Provider(nil, pathMocks))
}
