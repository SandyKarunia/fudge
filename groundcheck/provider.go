package groundcheck

import "gitlab.com/sandykarunia/fudge/utils"

// Provider ...
func Provider(sysUtils utils.System) GroundCheck {
	groundCheckOnce.Do(func() {
		groundCheck = &groundCheckImpl{
			sysUtils: sysUtils,
		}
	})
	return groundCheck
}
