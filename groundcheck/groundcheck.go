package groundcheck

import (
	"errors"
	"gitlab.com/sandykarunia/fudge/groundcheck/checkers"
)

var (
	errCheckAllFailed = errors.New("groundcheck: at least one check failed")
)

// GroundCheck is an entity that checks the machines where the program will run
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

	// we are splitting all the ifs because we want all checks to always run as they will provide nice messages

	if !g.c.CheckSudo() {
		errRes = errCheckAllFailed
	}

	if !g.c.CheckLibcapDevPkg() {
		errRes = errCheckAllFailed
	}

	if !g.c.CheckIsolateBinaryExists() {
		errRes = errCheckAllFailed
	}

	if !g.c.CheckIsolateBinaryExecutable() {
		errRes = errCheckAllFailed
	}

	return errRes
}
