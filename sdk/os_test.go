package sdk

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOsFunctionsImpl_Create(t *testing.T) {
	// mock osCreate
	originalOSCreate := osCreate
	osCreate = func(name string) (file *os.File, e error) {
		return nil, errors.New("name: " + name)
	}
	defer func() {
		osCreate = originalOSCreate
	}()

	obj := osFunctionsImpl{}
	res, err := obj.Create("param")
	assert.Nil(t, res)
	assert.Equal(t, errors.New("name: param"), err)
}

func TestOsFunctionsImpl_Getenv(t *testing.T) {
	// mock osGetenv
	originalOSGetenv := osGetenv
	osGetenv = func(key string) string {
		return key
	}
	defer func() {
		osGetenv = originalOSGetenv
	}()

	obj := osFunctionsImpl{}
	res := obj.Getenv("key1")
	assert.Equal(t, "key1", res)
}

func TestOsFunctionsImpl_Geteuid(t *testing.T) {
	// mock osGeteuid
	originalOSGeteuid := osGeteuid
	osGeteuid = func() int {
		return -123
	}
	defer func() {
		osGeteuid = originalOSGeteuid
	}()

	obj := osFunctionsImpl{}
	res := obj.Geteuid()
	assert.Equal(t, -123, res)
}

func TestOsFunctionsImpl_Open(t *testing.T) {
	// mock osOpen
	originalOSOpen := osOpen
	osOpen = func(name string) (file *os.File, e error) {
		return nil, errors.New("name: " + name)
	}
	defer func() {
		osOpen = originalOSOpen
	}()

	obj := osFunctionsImpl{}
	res, err := obj.Open("name1")
	assert.Nil(t, res)
	assert.Equal(t, errors.New("name: name1"), err)
}

func TestProvideOSFunctions(t *testing.T) {
	assert.Implements(t, (*OSFunctions)(nil), ProvideOSFunctions())
}
