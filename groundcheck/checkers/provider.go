package checkers

import "gitlab.com/sandykarunia/fudge/utils"

// Provider ...
func Provider(sysUtils utils.System, fileUtils utils.File, pathUtils utils.Path) Checkers {
	return &checkersImpl{
		sysUtils:  sysUtils,
		fileUtils: fileUtils,
		pathUtils: pathUtils,
	}
}
