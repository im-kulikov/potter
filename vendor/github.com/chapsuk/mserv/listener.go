package mserv

import (
	"context"
	"time"
)

type (
	// Serve wrapper for Listener interface
	Serve struct {
		shutdownTimeout time.Duration
		shutdownError   error
		server          Listener
	}

	Listener interface {
		ListenAndServe() error
		Shutdown(context.Context) error
	}
)

// NewListener returns new Listener wrapper
func NewListener(shutdownTimeout time.Duration, shutdownError error, server Listener) Server {
	if server == nil {
		log.Print("missing Listener, skip")
		return nil
	}

	return &Serve{
		shutdownTimeout: shutdownTimeout,
		shutdownError:   shutdownError,
		server:          server,
	}
}

// Start listener in goroutine
// write fatal msg by log if cant start server
func (s *Serve) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			// to ignore shutdown error
			if err != s.shutdownError {
				log.Fatalf("start listener error: %s", err)
			}
		}
	}()
}

// Stop listener with timeout
func (s *Serve) Stop() {
	if s == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Printf("stop listener error: %s", err)
	}
}
