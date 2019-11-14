package groundcheck

import (
	"errors"
	"github.com/sandykarunia/fudge/groundcheck/checkers"
)

var (
	errCheckAllFailed = errors.New("groundcheck: at least one check failed")
)

// GroundCheck is an entity that checks the machines where the program will run
// If it can solve the issue automatically, it will try to solve it
//go:generate mockery -name=GroundCheck
type GroundCheck interface {
	// CheckAll observes the environment / all necessities to run fudge program
	CheckAll() error
}

type groundCheckImpl struct {
	c checkers.Checkers
}

func (g *groundCheckImpl) CheckAll() error {
	var errRes error

	var checkerFuncs = []func() bool{
		g.c.CheckSudo,
		g.c.CheckLibcapDevPkg,
		g.c.CheckIsolateBinaryExists,
		g.c.CheckIsolateBinaryExecutable,
	}

	// we don't want to interrupt the checks (i.e. put return inside the loop)
	// because we want the loop to keep going, as the functions will provide nice messages

	for _, fn := range checkerFuncs {
		if !fn() {
			errRes = errCheckAllFailed
		}
	}

	return errRes
}
