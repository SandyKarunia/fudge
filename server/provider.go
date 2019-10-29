package server

import (
	"github.com/sandykarunia/fudge/groundcheck"
)

// Provider ...
func Provider(gc groundcheck.GroundCheck) Server {
	return &serverImpl{
		groundCheck: gc,
	}
}
