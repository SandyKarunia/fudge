package server

import (
	"gitlab.com/sandykarunia/fudge/groundcheck"
)

// Provider ...
func Provider(gc groundcheck.GroundCheck) Server {
	return &serverImpl{
		groundCheck: gc,
	}
}
