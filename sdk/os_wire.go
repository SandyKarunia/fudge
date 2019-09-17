//+build wireinject

package sdk

import "github.com/google/wire"

func OSInstance() OSFunctions {
	wire.Build(ProvideOSFunctions)
	return &osFunctionsImpl{}
}
