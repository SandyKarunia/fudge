//+build wireinject

package utils

import (
	"github.com/google/wire"
	"gitlab.com/sandykarunia/fudge/sdk"
)

func SystemInstance() System {
	wire.Build(ProvideSystem, sdk.Providers)
	return &systemImpl{}
}
