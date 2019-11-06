package handler

import "github.com/sandykarunia/fudge/grader"

// Provider ...
func Provider(grader grader.Grader) Handler {
	return &handlerImpl{
		grader: grader,
	}
}
