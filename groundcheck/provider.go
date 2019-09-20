package groundcheck

import "gitlab.com/sandykarunia/fudge/utils"

// Provider ...
func Provider(sysUtils utils.System) GroundCheck {
	return &groundCheckImpl{
		sysUtils: sysUtils,
	}
}
