package server

import (
	"context"
	"encoding/json"
	"net/http"

	"ninja_v1/internal/game"
	data "ninja_v1/internal/game/data"
)

type ActorsHandler struct {
	mux *http.ServeMux
}

func NewActorsHandler(ctx context.Context) *ActorsHandler {
	handler := &ActorsHandler{
		mux: http.NewServeMux(),
	}

	return handler
}

func (ah *ActorsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.mux.ServeHTTP(w, r)
}

func (ah *ActorsHandler) MakeGame() game.Game {
	return game.Game{
		Actors:    data.GetAllActors(),
		Modifiers: []game.ModifierTransaction{},
	}
}

func (ah *ActorsHandler) HandleGetActors(w http.ResponseWriter, r *http.Request) {
	g := ah.MakeGame()

	actorModifiers := game.GetActorModifiers(g)
	resolved := make([]game.ResolvedActor, 0, len(g.Actors))

	for _, a := range g.Actors {
		resolvedActor := game.ResolveActor(a, g.Modifiers, actorModifiers)
		resolved = append(resolved, resolvedActor)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resolved); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
