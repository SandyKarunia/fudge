package sdk

import "flag"

var (
	flagBool  = flag.Bool
	flagParse = flag.Parse
)

// FlagFunctions is an interface that represents flag library in golang sdk
//go:generate mockery -name=FlagFunctions
type FlagFunctions interface {
	Bool(name string, value bool, usage string) *bool
	Parse()
}

type flagFunctionsImpl struct{}

func (f *flagFunctionsImpl) Bool(name string, value bool, usage string) *bool {
	return flagBool(name, value, usage)
}

func (f *flagFunctionsImpl) Parse() {
	flagParse()
}

// ProvideFlagFunctions ...
func ProvideFlagFunctions() FlagFunctions {
	return &flagFunctionsImpl{}
}
