package server

import (
	"context"
	"log/slog"
	"net/http"
)

type Server struct {
	*http.Server
	logger *slog.Logger
}

func NewServer(ctx context.Context) *Server {
	logger := slog.Default()

	actorsHandler := NewActorsHandler(ctx)
	instancesHandler := NewInstancesHandler(ctx)

	mux := http.NewServeMux()
	api := http.NewServeMux()

	api.HandleFunc("GET /actors", actorsHandler.HandleGetActors)
	api.HandleFunc("GET /instances", instancesHandler.HandleGetGames)

	mux.Handle("/api/", http.StripPrefix("/api", api))
	mux.Handle("/games/", http.StripPrefix("/games", instancesHandler))

	return &Server{
		Server: &http.Server{
			Addr:    ":3005",
			Handler: mux,
		},
		logger: logger,
	}
}

func (s *Server) Run() error {
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
