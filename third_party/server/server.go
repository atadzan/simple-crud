package server

import (
	"context"

	"github.com/go-errors/errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Run - runs HTTP server
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              port,
		Handler:           handler,
		MaxHeaderBytes:    5 << 20,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return errors.New(err)
	}

	return nil
}

// GracefulShutDown - shutdowns HTTP server
func (s *Server) GracefulShutDown(ctx context.Context) error {
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(timeoutCtx); err != nil {
		return errors.New(err)
	}
	return nil
}
