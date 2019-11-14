package sandbox

import (
	"github.com/sandykarunia/fudge/logger/mocks"
	sdkmocks "github.com/sandykarunia/fudge/sdk/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProvider(t *testing.T) {
	mockSDKos := &sdkmocks.OSFunctions{}
	mockSDKos.On("Getenv", mock.Anything).Return("")
	mockLogger := &mocks.Logger{}
	mockLogger.On("Warn", mock.Anything).Return()
	mockLogger.On("Info", mock.Anything).Return()
	assert.Implements(t, (*Factory)(nil), Provider(mockSDKos, nil, nil, nil, mockLogger))
}
