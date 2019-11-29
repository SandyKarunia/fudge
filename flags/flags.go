package flags

import "github.com/sandykarunia/fudge/sdk"

// FlagName ...
type FlagName string

// Flags is an interface to get the command line flag values
//go:generate mockery -name=Flags
type Flags interface {
	// GetBool ...
	GetBool(flagName FlagName) bool
}

type flagsImpl struct {
	sdkFlag sdk.FlagFunctions

	// variable to store all the flag values
	boolValues map[FlagName]*bool
}

func (f *flagsImpl) GetBool(flagName FlagName) bool {
	if f.boolValues[flagName] == nil {
		return false
	}
	return *f.boolValues[flagName]
}

// Constants for all available flags
const (
	FakeSandbox FlagName = "fake-sandbox"
)

func (f *flagsImpl) init() {
	f.boolValues = make(map[FlagName]*bool)

	// list all the available flags
	f.boolValues[FakeSandbox] = f.sdkFlag.Bool(string(FakeSandbox), false, "indicates whether to use real (isolate) sandbox, or a fake sandbox")

	f.sdkFlag.Parse()
}
