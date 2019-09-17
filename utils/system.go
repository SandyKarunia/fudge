package utils

import (
	"gitlab.com/sandykarunia/fudge/sdk"
	"sync"
)

var systemOnce sync.Once
var system System

// System ...
//go:generate mockery -name=System
type System interface {
	// IsSudo returns true if this program run with sudo
	IsSudo() bool
}

type systemImpl struct {
	os sdk.OSFunctions
}

func (o *systemImpl) IsSudo() bool {
	if o.os.Geteuid() != 0 || len(o.os.Getenv("SUDO_UID")) == 0 ||
		len(o.os.Getenv("SUDO_GID")) == 0 || len(o.os.Getenv("SUDO_USER")) == 0 {
		return false
	}

	return true
}

// ProvideSystem ...
func ProvideSystem(os sdk.OSFunctions) System {
	systemOnce.Do(func() {
		system = &systemImpl{}
	})
	return system
}
