package sniffers

import "github.com/sandykarunia/fudge/utils"

// Provider ...
func Provider(sysUtils utils.System) Sniffers {
	return &sniffersImpl{
		sysUtils: sysUtils,
	}
}
