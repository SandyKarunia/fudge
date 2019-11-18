package sandbox

import (
	"github.com/sandykarunia/fudge/logger/mocks"
	utilsmocks "github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProvider(t *testing.T) {
	mockSystem := &utilsmocks.System{}
	mockSystem.On("IsControlGroupSupported").Return(true)
	mockLogger := &mocks.Logger{}
	mockLogger.On("Warn", mock.Anything).Return()
	mockLogger.On("Info", mock.Anything).Return()
	assert.Implements(t, (*Factory)(nil), Provider(nil, nil, nil, mockSystem, mockLogger))
}
