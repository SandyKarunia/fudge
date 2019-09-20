package groundcheck

import (
	"errors"
	"gitlab.com/sandykarunia/fudge/utils"
)

// GroundCheck is an entity that checks the machines where the program will run
//go:generate mockery -name=GroundCheck
type GroundCheck interface {
	// CheckAll observes the environment / all necessities to run fudge program
	// if the environment is not suitable to run fudge, it will return an error which contains the error message
	CheckAll() error
}

type groundCheckImpl struct {
	sysUtils utils.System
}

func (g *groundCheckImpl) CheckAll() error {
	// Check whether the we are in sudo environment or not
	if !g.sysUtils.IsSudo() {
		return errors.New("please run this program as root, we need root to run the isolate binary")
	}
	return nil
}
