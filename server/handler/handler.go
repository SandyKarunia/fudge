package handler

import (
	"github.com/sandykarunia/fudge/grader"
	"github.com/sandykarunia/fudge/logger"
	"net/http"
)

// Handler is an interface for all handlers for requests
type Handler interface {
	Grade(w http.ResponseWriter, r *http.Request)
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	grader grader.Grader
	logger logger.Logger
}
