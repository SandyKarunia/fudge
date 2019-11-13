package sandbox

import (
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
)

// Provider ...
func Provider(sdkOS sdk.OSFunctions, sdkIO sdk.IOFunctions, path utils.Path, system utils.System) Factory {
	return &factoryImpl{
		sdkOS:       sdkOS,
		sdkIO:       sdkIO,
		utilsPath:   path,
		utilsSystem: system,
	}
}
