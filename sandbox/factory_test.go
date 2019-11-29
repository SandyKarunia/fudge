package sandbox

import (
	flagsMocks "github.com/sandykarunia/fudge/flags/mocks"
	utilsMocks "github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"testing"
	"time"
)

func TestFactoryImpl_NewPreparedSandbox(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	mockPath := &utilsMocks.Path{}
	mockPath.On("IsolateBinary").Return("")
	mockSystem := &utilsMocks.System{}
	mockSystem.On("Execute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	mockFlags := &flagsMocks.Flags{}
	mockFlags.On("GetBool", mock.Anything).Return(false)
	obj := factoryImpl{utilsPath: mockPath, utilsSystem: mockSystem, flags: mockFlags}

	// should return different IDs most of the time
	usedIDs := map[uint32]bool{}
	for i := 0; i < 5; i++ {
		sb, _ := obj.NewPreparedSandbox()
		sbID := sb.GetID()

		assert.NotContains(t, usedIDs, sbID, "usedIDs map contains duplicate ID %d", sbID)
		usedIDs[sbID] = true
	}
}
