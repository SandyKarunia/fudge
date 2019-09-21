package utils

import (
	"errors"
	"gitlab.com/sandykarunia/fudge/sdk"
	"strings"
)

// System ...
//go:generate mockery -name=System
type System interface {
	// IsSudo returns true if this program run with sudo
	IsSudo() bool

	// VerifyPkgInstalled returns error message if the specified package is NOT installed
	VerifyPkgInstalled(pkgName string) error

	// GetFudgeDir returns config directory for all fudge related stuff, always ends with "/"
	GetFudgeDir() string
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

func (o *systemImpl) GetFudgeDir() string {
	homeDir, err := o.os.UserHomeDir()
	if err != nil {
		// TODO log error
	}

	// if it doesn't end with "/", add it
	if !strings.HasSuffix(homeDir, "/") {
		homeDir += "/"
	}

	return homeDir + ".fudge/"
}

// ProvideSystem ...
func ProvideSystem(os sdk.OSFunctions, exec sdk.ExecFunctions) System {
	return &systemImpl{
		os:   os,
		exec: exec,
	}
}
