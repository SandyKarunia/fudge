package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sandykarunia/fudge/groundcheck"
	"github.com/sandykarunia/fudge/logger"
	"github.com/sandykarunia/fudge/server/handler"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server is an interface for a fudge server
type Server interface {
	// Start starts fudge server
	Start()
}

type serverImpl struct {
	groundCheck groundcheck.GroundCheck
	handler     handler.Handler
	logger      logger.Logger
}

func (s *serverImpl) Start() {
	// before we do something, do ground check first
	if err := s.groundCheck.CheckAll(); err != nil {
		os.Exit(1)
	}
	// before we do something, sniff all to give user the information about the environment
	s.groundCheck.SniffAll()

	defaultPort := 8080

	r := mux.NewRouter()
	r.HandleFunc("/health_check", s.handler.HealthCheck)
	r.HandleFunc("/grade", s.handler.Grade).Methods(http.MethodPost)

	addr := fmt.Sprintf("0.0.0.0:%d", defaultPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		s.logger.Info("Server started at %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("Failed to listen and serve, err: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, we could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if our application should wait for other services
	// to finalize based on context cancellation.
	s.logger.Info("Shutting down")
	s.logger.FlushBuffer()
	os.Exit(0)
}
