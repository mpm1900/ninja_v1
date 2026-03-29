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

func (dh *DataHandler) HandleGetActors(w http.ResponseWriter, r *http.Request) {
	actors := slices.Collect(maps.Values(data.ACTORS))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(actors); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (dh *DataHandler) HandleGetActions(w http.ResponseWriter, r *http.Request) {
	actions := slices.Collect(maps.Values(data.ACTIONS))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(actions); err != nil {
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

func (dh *DataHandler) HandleGetTriggerTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(game.TRIGGERS); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (dh *DataHandler) HandleIsActionContextValid(w http.ResponseWriter, r *http.Request) {
	var context game.Context
	err := json.NewDecoder(r.Body).Decode(&context)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if context.ActionID == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	action, ok := data.ACTIONS[*context.ActionID]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	valid := action.ContextValidate(context)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(valid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
