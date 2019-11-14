package groundcheck

import (
	"errors"
	"github.com/fatih/color"
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

type checkerAndSolver struct {
	checkerFunc func() bool
	solverFunc  func()
}

func (g *groundCheckImpl) CheckAll() error {
	var errRes error

	var checkerFuncs = []checkerAndSolver{
		{g.c.CheckSudo, nil},
		{g.c.CheckLibcapDevPkg, nil},
		{g.c.CheckIsolateBinaryExists, nil},
		{g.c.CheckIsolateBinaryExecutable, nil},
	}

	// we don't want to interrupt the checks (i.e. put return inside the loop)
	// because we want the loop to keep going, as the functions will provide nice messages

	for _, cns := range checkerFuncs {
		// try to check
		ok := cns.checkerFunc()
		if ok {
			continue
		}

		// if there is no solver function, then error and just continue
		if cns.solverFunc == nil {
			errRes = errCheckAllFailed
			continue
		}

		color.HiMagenta("Trying to solve the problem...")

		// try to solve
		cns.solverFunc()

		// check again
		ok = cns.checkerFunc()
		if !ok {
			errRes = errCheckAllFailed
		}
	}

	return errRes
}
