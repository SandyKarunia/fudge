//+build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/sandykarunia/fudge/groundcheck"
	"github.com/sandykarunia/fudge/groundcheck/checkers"
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
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
