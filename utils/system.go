package utils

import (
	"errors"
	"github.com/sandykarunia/fudge/sdk"
)

// System ...
//go:generate mockery -name=System
type System interface {
	// IsSudo returns true if this program run with sudo
	IsSudo() bool

	// VerifyPkgInstalled returns error message if the specified package is NOT installed
	VerifyPkgInstalled(pkgName string) error

	// Execute executes a command, and return the output (stdout + stderr) and error
	Execute(cmd string, args ...string) (string, error)

	// GetHMACSecretFromEnv returns a secret string for HMAC authentication from environment variable
	GetHMACSecretFromEnv() string

	// IsControlGroupSuppoerted returns true if control group is supported in current machine
	// by checking CONFIG_CPUSETS environment variable value
	IsControlGroupSupported() bool
}

type systemImpl struct {
	os   sdk.OSFunctions
	exec sdk.ExecFunctions
}

func (o *systemImpl) IsSudo() bool {
	if o.os.Geteuid() != 0 || len(o.os.Getenv("SUDO_UID")) == 0 ||
		len(o.os.Getenv("SUDO_GID")) == 0 || len(o.os.Getenv("SUDO_USER")) == 0 {
		return false
	}

	return true
}

func (o *systemImpl) VerifyPkgInstalled(pkgName string) error {
	cmd := o.exec.Command("dpkg", "-s", pkgName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	return nil
}

func (o *systemImpl) Execute(name string, args ...string) (string, error) {
	cmd := o.exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func (o *systemImpl) GetHMACSecretFromEnv() string {
	return o.os.Getenv("FUDGE_HMAC")
}

func (o *systemImpl) IsControlGroupSupported() bool {
	return len(o.os.Getenv("CONFIG_CPUSETS")) > 0
}

// ProvideSystem ...
func ProvideSystem(os sdk.OSFunctions, exec sdk.ExecFunctions) System {
	return &systemImpl{
		os:   os,
		exec: exec,
	}
}
