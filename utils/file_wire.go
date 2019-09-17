//+build wireinject

package utils

import (
	"github.com/google/wire"
	"gitlab.com/sandykarunia/fudge/sdk"
)

func FileInstance() File {
	wire.Build(ProvideFile, sdk.Providers)
	return &fileImpl{}
}
