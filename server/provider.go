package server

import (
	"github.com/sandykarunia/fudge/groundcheck"
	"github.com/sandykarunia/fudge/server/handler"
)

// Provider ...
func Provider(gc groundcheck.GroundCheck, handler handler.Handler) Server {
	return &serverImpl{
		groundCheck: gc,
		handler:     handler,
	}
}
