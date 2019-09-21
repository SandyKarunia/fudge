package groundcheck

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/sandykarunia/fudge/groundcheck/checkers/mocks"
	"testing"
)

func TestGroundCheckImpl_CheckAll(t *testing.T) {
	tests := []struct {
		desc               string
		checkSudo          bool
		checkLibcapDevPkg  bool
		checkIsolateBinary bool
		want               error
	}{
		{
			desc:               "check sudo is false",
			checkSudo:          false,
			checkLibcapDevPkg:  true,
			checkIsolateBinary: true,
			want:               errCheckAllFailed,
		},
		{
			desc:               "check sudo is true, rest is false",
			checkSudo:          true,
			checkLibcapDevPkg:  false,
			checkIsolateBinary: false,
			want:               errCheckAllFailed,
		},
		{
			desc:               "all true",
			checkSudo:          true,
			checkLibcapDevPkg:  true,
			checkIsolateBinary: true,
			want:               nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mockCheckers := &mocks.Checkers{}

			mockCheckers.On("CheckSudo").Return(test.checkSudo)
			mockCheckers.On("CheckLibcapDevPkg").Return(test.checkLibcapDevPkg)
			mockCheckers.On("CheckIsolateBinary").Return(test.checkIsolateBinary)

			obj := &groundCheckImpl{c: mockCheckers}
			res := obj.CheckAll()
			assert.Equal(t, test.want, res, test.desc)
		})
	}
}
