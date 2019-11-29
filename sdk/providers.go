package sdk

import "github.com/google/wire"

// Providers contain all providers from current package
var Providers = wire.NewSet(
	ProvideFlagFunctions,
	ProvideIOFunctions,
	ProvideOSFunctions,
	ProvideHTTPFunctions,
	ProvideExecFunctions)
