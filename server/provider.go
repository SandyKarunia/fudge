package server

import (
	"github.com/sandykarunia/fudge/groundcheck"
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/server/handler"
)

// Provider ...
func Provider(gc groundcheck.GroundCheck, handler handler.Handler, logger logger.Logger) Server {
	return &serverImpl{
		groundCheck: gc,
		handler:     handler,
		logger:      logger,
	}
}
