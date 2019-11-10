package logger

import (
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
)

// Provider ...
func Provider(os sdk.OSFunctions, sys utils.System) Logger {
	obj := &loggerImpl{
		os:  os,
		sys: sys,
	}
	return obj.init()
}
