//+build wireinject

package sdk

import "github.com/google/wire"

func IOInstance() IOFunctions {
	wire.Build(ProvideIOFunctions)
	return &ioFunctionsImpl{}
}
