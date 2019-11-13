package sandbox

import (
	"github.com/sandykarunia/fudge/utils/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactoryImpl_NewSandbox(t *testing.T) {
	mockPath := &mocks.Path{}
	obj := factoryImpl{utilsPath: mockPath}

	// should return different IDs most of the time
	usedIDs := map[int]bool{}
	for i := 0; i < 100; i++ {
		sb := obj.NewSandbox()
		sbID := sb.GetID()

		assert.NotContains(t, usedIDs, sbID, "usedIDs map contains duplicate ID %d", sbID)
		usedIDs[sbID] = true
	}
}
