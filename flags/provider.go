package flags

import "github.com/sandykarunia/fudge/sdk"

// Provider ...
func Provider(sdkFlag sdk.FlagFunctions) Flags {
	obj := &flagsImpl{
		sdkFlag: sdkFlag,
	}
	obj.init()

	return obj
}
