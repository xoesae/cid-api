package http

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type Config struct {
	Port   string
	Router *chi.Mux
}

type Server struct {
	server *http.Server
	config Config
}

func New(c Config) *Server {
	_server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", c.Port),
		Handler: c.Router,
	}

	return &Server{
		server: _server,
		config: c,
	}
}

func (s *Server) Start() error {
	slog.Info("starting server", "port", s.config.Port)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
