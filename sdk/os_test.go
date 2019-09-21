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

func TestOsFunctionsImpl_UserHomeDir(t *testing.T) {
	// mock osUserHomeDir
	originalOSUserHomeDir := osUserHomeDir
	osUserHomeDir = func() (s string, e error) {
		return "home", nil
	}
	defer func() {
		osUserHomeDir = originalOSUserHomeDir
	}()

	obj := osFunctionsImpl{}
	res, err := obj.UserHomeDir()
	assert.Nil(t, err)
	assert.Equal(t, "home", res)
}

func TestOsFunctionsImpl_Stat(t *testing.T) {
	// mock osStat
	originalOSStat := osStat
	osStat = func(name string) (info os.FileInfo, e error) {
		return nil, errors.New(name)
	}
	defer func() {
		osStat = originalOSStat
	}()

	obj := osFunctionsImpl{}
	res, err := obj.Stat("woi")
	assert.Nil(t, res)
	assert.Equal(t, errors.New("woi"), err)
}

func TestOsFunctionsImpl_IsNotExist(t *testing.T) {
	// mock osIsNotExist
	originalOSIsNotExist := osIsNotExist
	osIsNotExist = func(err error) bool {
		return true
	}
	defer func() {
		osIsNotExist = originalOSIsNotExist
	}()

	obj := osFunctionsImpl{}
	res := obj.IsNotExist(nil)
	assert.True(t, res)
}

func TestProvideOSFunctions(t *testing.T) {
	assert.Implements(t, (*OSFunctions)(nil), ProvideOSFunctions())
}
