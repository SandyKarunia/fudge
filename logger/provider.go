package logger

import "github.com/sandykarunia/fudge/sdk"

// ProvideStdLogger ...
func ProvideStdLogger(fmt sdk.FmtFunctions) Logger {
	return &stdLogger{
		fmt: fmt,
	}
}
