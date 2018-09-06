package mserv

import (
	"context"
	"net/http"
	"time"
)

// HTTPServer wrapper of http.Server
type HTTPServer struct {
	shutdownTimeout time.Duration
	server          *http.Server
}

// NewHTTPServer returns new http.Server wrapper
func NewHTTPServer(shutdownTimeout time.Duration, server *http.Server) Server {
	if server == nil {
		log.Print("missing http.Server, skip")
		return nil
	}

	if server.Addr == "" {
		log.Print("missing bind address for http.Server, skip")
		return nil
	}

	return &HTTPServer{
		shutdownTimeout: shutdownTimeout,
		server:          server,
	}
}

// Start http server in goroutine
// write fatal msg by log if cant start server
func (h *HTTPServer) Start() {
	if h == nil {
		return
	}

	go func() {
		if err := h.server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("start http server error: %s", err)
			}
		}
	}()
}

// Stop http server with timeout
func (h *HTTPServer) Stop() {
	if h == nil {
		return
	}

	if h.shutdownTimeout == 0 {
		if err := h.server.Close(); err != nil {
			log.Printf("stop http server error: %s", err)
		}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.shutdownTimeout)
	defer cancel()

	if err := h.server.Shutdown(ctx); err != nil {
		log.Printf("stop http server error: %s", err)
	}
}
