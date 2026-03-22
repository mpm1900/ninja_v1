package server

import (
	"context"
	"encoding/json"
	"maps"
	"net/http"
	"slices"

	"ninja_v1/internal/game"
	data "ninja_v1/internal/game/data"
)

type DataHandler struct {
	mux *http.ServeMux
}

func NewDataHandler(ctx context.Context) *DataHandler {
	handler := &DataHandler{
		mux: http.NewServeMux(),
	}

	return handler
}

func (dh *DataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dh.mux.ServeHTTP(w, r)
}

func (dh *DataHandler) MakeGame() game.Game {
	return game.Game{
		Actors:    data.GetAllActors(),
		Modifiers: []game.ModifierTransaction{},
	}
}

func (dh *DataHandler) HandleGetActors(w http.ResponseWriter, r *http.Request) {
	g := dh.MakeGame()
	resolved := make([]game.ResolvedActor, 0, len(g.Actors))

	for _, a := range g.Actors {
		resolvedActor := game.ResolveActor(a, g)
		resolved = append(resolved, resolvedActor)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resolved); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (dh *DataHandler) HandleGetModifiers(w http.ResponseWriter, r *http.Request) {
	modifiers := slices.Collect(maps.Values(data.MODIFIERS))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(modifiers); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
