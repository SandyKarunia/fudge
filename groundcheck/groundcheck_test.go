package groundcheck

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/sandykarunia/fudge/utils/mocks"
	"testing"
)

func TestGroundCheckImpl_CheckAll(t *testing.T) {
	tests := []struct {
		desc   string
		isSudo bool
		want   error
	}{
		{
			desc:   "It's in sudo environment",
			isSudo: true,
			want:   nil,
		},
		{
			desc:   "It's not in sudo environment",
			isSudo: false,
			want:   errNotSudo,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mockSysUtils := &mocks.System{}
			gc := &groundCheckImpl{sysUtils: mockSysUtils}

			mockSysUtils.On("IsSudo").Return(test.isSudo)

			res := gc.CheckAll()
			assert.Equal(t, test.want, res, test.want)
		})
	}
}
