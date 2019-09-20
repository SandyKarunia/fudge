package sdk

import (
	"os"
)

var (
	osOpen    = os.Open
	osCreate  = os.Create
	osGeteuid = os.Geteuid
	osGetenv  = os.Getenv
)

// OSFunctions is an interface that represents os library in golang sdk
//go:generate mockery -name=OSFunctions
type OSFunctions interface {
	Open(name string) (*os.File, error)
	Create(name string) (*os.File, error)
	Geteuid() int
	Getenv(key string) string
}

type osFunctionsImpl struct{}

func (o *osFunctionsImpl) Open(name string) (*os.File, error) {
	return osOpen(name)
}

func (o *osFunctionsImpl) Create(name string) (*os.File, error) {
	return osCreate(name)
}

func (o *osFunctionsImpl) Geteuid() int {
	return osGeteuid()
}

func (o *osFunctionsImpl) Getenv(key string) string {
	return osGetenv(key)
}

// ProvideOSFunctions ...
func ProvideOSFunctions() OSFunctions {
	return &osFunctionsImpl{}
}
