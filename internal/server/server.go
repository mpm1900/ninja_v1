package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"

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
	api.HandleFunc("GET /items", dataHandler.HandleGetItems)

	api.HandleFunc("GET /instances", instancesHandler.HandleGetGames)

	mux.Handle("/api/", http.StripPrefix("/api", api))
	mux.Handle("/socket/", http.StripPrefix("/socket", authenticatedInstancesHandler))

	// CORS and other global middleware
	handler := http.Handler(mux)
	handler = withCORS(handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}
	if port[0] != ':' {
		port = ":" + port
	}

	return &Server{
		Server: &http.Server{
			Addr:    port,
			Handler: handler,
		},
		logger: logger,
	}
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) Run() error {
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
