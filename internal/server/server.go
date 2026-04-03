package server

import (
	"context"
	"log/slog"
	"net/http"

	"ninja_v1/internal/auth"
	"ninja_v1/internal/db"
)

type Server struct {
	*http.Server
	logger *slog.Logger
}

func NewServer(ctx context.Context, queries *db.Queries) *Server {
	logger := slog.Default()

	dataHandler := NewDataHandler(ctx)
	instancesHandler := NewInstancesHandler(ctx)
	authenticatedInstancesHandler := auth.WithSession(instancesHandler.ServeHTTP, queries)

	mux := http.NewServeMux()
	api := http.NewServeMux()

	api.HandleFunc("POST /auth/signup", handleSignUp(ctx, queries))
	api.HandleFunc("POST /auth/login", handleLogin(ctx, queries))
	api.HandleFunc("POST /auth/logout", auth.WithSession(handleLogout(ctx, queries), queries))
	api.HandleFunc("GET  /auth/me", auth.WithSession(handleMe(), queries))

	api.HandleFunc("GET /actions", dataHandler.HandleGetActions)
	api.HandleFunc("GET /actors", dataHandler.HandleGetActors)

	api.HandleFunc("GET /instances", instancesHandler.HandleGetGames)
	api.HandleFunc("POST /{instanceID}/targets", instancesHandler.HandleGetTargets)

	mux.Handle("/api/", http.StripPrefix("/api", api))
	mux.Handle("/socket/", http.StripPrefix("/socket", authenticatedInstancesHandler))

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
