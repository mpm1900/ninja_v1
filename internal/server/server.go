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

	dataHandler := NewDataHandler(ctx)
	instancesHandler := NewInstancesHandler(ctx)

	mux := http.NewServeMux()
	api := http.NewServeMux()

	api.HandleFunc("GET /actions", dataHandler.HandleGetActions)
	api.HandleFunc("GET /actors", dataHandler.HandleGetActors)
	api.HandleFunc("GET /modifiers", dataHandler.HandleGetModifiers)
	api.HandleFunc("GET /triggers", dataHandler.HandleGetTriggerTypes)
	api.HandleFunc("POST /{actionID}/validate", dataHandler.HandleIsActionContextValid)

	api.HandleFunc("GET /instances", instancesHandler.HandleGetGames)
	api.HandleFunc("POST /{instanceID}/{actionID}/targets", instancesHandler.HandleGetTargets)

	mux.Handle("/api/", http.StripPrefix("/api", api))
	mux.Handle("/socket/", http.StripPrefix("/socket", instancesHandler))

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
