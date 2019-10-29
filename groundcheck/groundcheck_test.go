package groundcheck

import (
	"github.com/stretchr/testify/assert"
	"github.com/sandykarunia/fudge/groundcheck/checkers/mocks"
	"testing"
)

func TestGroundCheckImpl_CheckAll(t *testing.T) {
	tests := []struct {
		desc                         string
		checkSudo                    bool
		checkLibcapDevPkg            bool
		checkIsolateBinaryExists     bool
		checkIsolateBinaryExecutable bool
		want                         error
	}{
		{
			desc:                         "check sudo is false",
			checkSudo:                    false,
			checkLibcapDevPkg:            true,
			checkIsolateBinaryExists:     true,
			checkIsolateBinaryExecutable: true,
			want:                         errCheckAllFailed,
		},
		{
			desc:                     "check sudo is true, rest is false",
			checkSudo:                true,
			checkLibcapDevPkg:        false,
			checkIsolateBinaryExists: false,
			want:                     errCheckAllFailed,
		},
		{
			desc:                         "all true",
			checkSudo:                    true,
			checkLibcapDevPkg:            true,
			checkIsolateBinaryExists:     true,
			checkIsolateBinaryExecutable: true,
			want:                         nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mockCheckers := &mocks.Checkers{}

			mockCheckers.On("CheckSudo").Return(test.checkSudo)
			mockCheckers.On("CheckLibcapDevPkg").Return(test.checkLibcapDevPkg)
			mockCheckers.On("CheckIsolateBinaryExists").Return(test.checkIsolateBinaryExists)
			mockCheckers.On("CheckIsolateBinaryExecutable").Return(test.checkIsolateBinaryExecutable)

			obj := &groundCheckImpl{c: mockCheckers}
			res := obj.CheckAll()
			assert.Equal(t, test.want, res, test.desc)
		})
	}
}
