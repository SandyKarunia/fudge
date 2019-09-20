//+build wireinject

package server

import (
	"github.com/google/wire"
	"gitlab.com/sandykarunia/fudge/groundcheck"
	"gitlab.com/sandykarunia/fudge/sdk"
	"gitlab.com/sandykarunia/fudge/utils"
)

func Instance() Server {
	wire.Build(
		Provider, groundcheck.Provider, utils.Providers, sdk.Providers,
	)
	return &serverImpl{}
}
