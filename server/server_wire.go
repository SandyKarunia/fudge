//+build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/sandykarunia/fudge/flags"
	"github.com/sandykarunia/fudge/grader"
	"github.com/sandykarunia/fudge/grader/taskrunner"
	"github.com/sandykarunia/fudge/groundcheck"
	"github.com/sandykarunia/fudge/groundcheck/checkers"
	"github.com/sandykarunia/fudge/groundcheck/sniffers"
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/sandbox"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/server/handler"
	"github.com/sandykarunia/fudge/utils"
)

func Instance() Server {
	wire.Build(
		Provider,
		checkers.Provider,
		flags.Provider,
		grader.Provider,
		groundcheck.Provider,
		handler.Provider,
		logger.Provider,
		sandbox.Provider,
		sdk.Providers,
		sniffers.Provider,
		taskrunner.Provider,
		utils.Providers,
	)
	return &serverImpl{}
}
