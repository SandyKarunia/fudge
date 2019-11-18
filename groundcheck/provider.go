package groundcheck

import (
	"github.com/sandykarunia/fudge/groundcheck/checkers"
	"github.com/sandykarunia/fudge/groundcheck/sniffers"
)

// Provider ...
func Provider(checkers checkers.Checkers, sniffers sniffers.Sniffers) GroundCheck {
	return &groundCheckImpl{
		c: checkers,
		s: sniffers,
	}
}
