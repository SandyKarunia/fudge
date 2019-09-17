//+build wireinject

package utils

import "github.com/google/wire"

func StringInstance() String {
	wire.Build(ProvideString)
	return &stringImpl{}
}
