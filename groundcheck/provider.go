package groundcheck

import "gitlab.com/sandykarunia/fudge/groundcheck/checkers"

// Provider ...
func Provider(checkers checkers.Checkers) GroundCheck {
	return &groundCheckImpl{
		c: checkers,
	}
}
