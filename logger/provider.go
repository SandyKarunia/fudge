package logger

import (
	"github.com/sandykarunia/fudge/sdk"
	"github.com/sandykarunia/fudge/utils"
)

// Provider ...
func Provider(os sdk.OSFunctions, path utils.Path) Logger {
	obj := &loggerImpl{
		os:   os,
		path: path,
	}
	return obj.init()
}
