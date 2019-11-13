package handler

import (
	"github.com/sandykarunia/fudge/grader"
	"github.com/sandykarunia/fudge/logger"
)

// Provider ...
func Provider(grader grader.Grader, logger logger.Logger) Handler {
	return &handlerImpl{
		grader: grader,
		logger: logger,
	}
}
