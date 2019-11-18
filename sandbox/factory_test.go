package sandbox

import (
	"github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"testing"
	"time"
)

func TestFactoryImpl_NewSandbox(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	mockPath := &mocks.Path{}
	mockPath.On("IsolateBinary").Return("")
	mockSystem := &mocks.System{}
	mockSystem.On("Execute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	obj := factoryImpl{utilsPath: mockPath, utilsSystem: mockSystem}

	// should return different IDs most of the time
	usedIDs := map[uint32]bool{}
	for i := 0; i < 5; i++ {
		sb, _ := obj.NewPreparedSandbox()
		sbID := sb.GetID()

		assert.NotContains(t, usedIDs, sbID, "usedIDs map contains duplicate ID %d", sbID)
		usedIDs[sbID] = true
	}
}
