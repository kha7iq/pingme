package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kha7iq/pingme/internal/handlers"
	"github.com/kha7iq/pingme/internal/middleware"
)

// Server represents the HTTP server
type Server struct {
	httpServer *http.Server
	host       string
	port       string
}

// New creates a new server instance
func New(host, port string) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

// Start initializes and starts the HTTP server with graceful shutdown
func (s *Server) Start() error {
	// Create router/mux
	mux := http.NewServeMux()

	// Setup routes
	s.setupRoutes(mux)

	// Apply middleware
	handler := s.applyMiddleware(mux)

	// Configure HTTP server
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	s.httpServer = &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to listen for errors from the server
	serverErrors := make(chan error, 1)

	// Start HTTP server in a goroutine
	go func() {
		log.Printf("Starting webhook server on %s", addr)
		log.Printf("POST webhooks to: http://%s/webhook", addr)
		serverErrors <- s.httpServer.ListenAndServe()
	}()

	// Channel to listen for interrupt signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// Block until we receive a signal or server error
	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server error: %w", err)
		}

	case sig := <-shutdown:
		log.Printf("Received signal %v, initiating graceful shutdown", sig)

		// Create context with timeout for shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Attempt graceful shutdown
		if err := s.httpServer.Shutdown(ctx); err != nil {
			log.Printf("Graceful shutdown failed, forcing close: %v", err)
			if closeErr := s.httpServer.Close(); closeErr != nil {
				return fmt.Errorf("force close error: %w", closeErr)
			}
			return fmt.Errorf("graceful shutdown error: %w", err)
		}

		log.Println("Server stopped gracefully")
	}

	return nil
}

// setupRoutes configures all HTTP routes
func (s *Server) setupRoutes(mux *http.ServeMux) {
	// Root info endpoint
	mux.HandleFunc("/", s.infoHandler)

	// Health check endpoint
	mux.HandleFunc("/health", s.healthHandler)

	// Webhook endpoint
	webhookHandler := handlers.NewWebhookHandler()
	mux.Handle("/webhook", webhookHandler)
}

// applyMiddleware wraps the handler with middleware chain
func (s *Server) applyMiddleware(handler http.Handler) http.Handler {
	// Apply middleware in reverse order (last defined = first executed)

	// Logging middleware (outermost - logs everything)
	handler = middleware.Logging(handler)

	// Authentication middleware (if enabled via env var)
	if os.Getenv("PINGME_AUTH_METHOD") != "" && os.Getenv("PINGME_AUTH_METHOD") != "none" {
		handler = middleware.Auth(handler)
	}

	// Recovery middleware (catches panics)
	handler = middleware.Recovery(handler)

	return handler
}

// infoHandler provides server information
func (s *Server) infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := `{
  "service": "PingMe Webhook Server",
  "endpoints": {
    "webhook": "/webhook (POST)",
    "health": "/health (GET)"
  },
  "usage": "Configure services via environment variables, then POST JSON to /webhook"
}`
	w.Write([]byte(response))
}

// healthHandler handles health check requests
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","service":"pingme"}`))
}
