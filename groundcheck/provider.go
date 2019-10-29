package groundcheck

import "github.com/sandykarunia/fudge/groundcheck/checkers"

// Provider ...
func Provider(checkers checkers.Checkers) GroundCheck {
	return &groundCheckImpl{
		c: checkers,
	}
}
