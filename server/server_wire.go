//+build wireinject

package server

import (
	"github.com/google/wire"
	"gitlab.com/sandykarunia/fudge/groundcheck"
	"gitlab.com/sandykarunia/fudge/groundcheck/checkers"
	"gitlab.com/sandykarunia/fudge/sdk"
	"gitlab.com/sandykarunia/fudge/utils"
)

func Instance() Server {
	wire.Build(
		Provider,
		checkers.Provider,
		groundcheck.Provider,
		sdk.Providers,
		utils.Providers,
	)
	return &serverImpl{}
}
