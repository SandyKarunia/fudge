package logger

import "github.com/sandykarunia/fudge/sdk"

// Provider ...
func Provider(os sdk.OSFunctions) Logger {
	return &loggerImpl{
		os: os,
	}
}
