package sandbox

import (
	"github.com/sandykarunia/fudge/flags"
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
)

// Provider ...
func Provider(
	sdkOS sdk.OSFunctions,
	sdkIO sdk.IOFunctions,
	flags flags.Flags,
	path utils.Path,
	system utils.System,
	logger logger.Logger) Factory {
	factory := &factoryImpl{
		sdkOS:       sdkOS,
		sdkIO:       sdkIO,
		flags:       flags,
		utilsPath:   path,
		utilsSystem: system,
		logger:      logger,
	}
	factory.init()

	return factory
}
