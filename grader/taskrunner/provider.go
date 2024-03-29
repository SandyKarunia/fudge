package taskrunner

import (
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sandbox"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
)

// Provider ...
func Provider(factory sandbox.Factory, logger logger.Logger, utilsString utils.String, sdkHTTP sdk.HTTPFunctions) TaskRunner {
	return &taskRunnerImpl{
		sbFactory:   factory,
		logger:      logger,
		utilsString: utilsString,
		sdkHTTP:     sdkHTTP,
	}
}
