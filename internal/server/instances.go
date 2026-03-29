package server

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"ninja_v1/internal/auth"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data"
	"ninja_v1/internal/instance"

	"github.com/google/uuid"
)

type InstancesHandler struct {
	mux         *http.ServeMux
	instances   map[uuid.UUID]*instance.Instance
	instancesMu sync.RWMutex
}

func NewInstancesHandler(ctx context.Context) *InstancesHandler {
	instanceID := uuid.New()
	defaultInstance := instance.NewInstance(ctx, instanceID)
	handler := &InstancesHandler{
		mux: http.NewServeMux(),
		instances: map[uuid.UUID]*instance.Instance{
			instanceID: defaultInstance,
		},
	}

	go defaultInstance.Run()

	handler.mux.HandleFunc("GET /{instanceID}/connect", handler.handleGameConnection(ctx))
	return handler
}

func (ih *InstancesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ih.mux.ServeHTTP(w, r)
}

func (ih *InstancesHandler) NewInstance(instanceID uuid.UUID, ctx context.Context) *instance.Instance {
	instance := instance.NewInstance(ctx, instanceID)
	ih.instancesMu.Lock()
	ih.instances[instance.ID] = instance
	ih.instancesMu.Unlock()
	go instance.Run()

	return instance
}

func (ih *InstancesHandler) GetInstance(instanceID uuid.UUID) (*instance.Instance, bool) {
	ih.instancesMu.RLock()
	defer ih.instancesMu.RUnlock()
	instance, ok := ih.instances[instanceID]

	return instance, ok
}

func (ih *InstancesHandler) GetAllInstances() []instance.Instance {
	ih.instancesMu.RLock()
	defer ih.instancesMu.RUnlock()
	games := make([]instance.Instance, 0, len(ih.instances))
	for _, instance := range ih.instances {
		games = append(games, *instance)
	}
	return games
}

func (ih *InstancesHandler) HandleGetGames(w http.ResponseWriter, r *http.Request) {
	games := ih.GetAllInstances()
	if len(games) == 0 {
		games = make([]instance.Instance, 0)
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(games)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ih *InstancesHandler) HandleGetTargets(w http.ResponseWriter, r *http.Request) {
	instanceID, err := uuid.Parse(r.PathValue("instanceID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var context game.Context
	err = json.NewDecoder(r.Body).Decode(&context)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if context.ActionID == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	instance, ok := ih.GetInstance(instanceID)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	action, ok := data.ACTIONS[*context.ActionID]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	targets := instance.Game.GetActors(func(a game.Actor) bool {
		return action.TargetPredicate(a, context)
	})
	targetIDs := make([]uuid.UUID, 0, len(targets))
	for _, a := range targets {
		targetIDs = append(targetIDs, a.ID)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(targetIDs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ih *InstancesHandler) handleGameConnection(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := auth.AuthenticatedUserFromContext(r.Context())
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		instanceID, err := uuid.Parse(r.PathValue("instanceID"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		i, ok := ih.GetInstance(instanceID)
		if !ok {
			i = ih.NewInstance(instanceID, ctx)
		}

		client := instance.NewClient(i)
		client.AttachUser(&game.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
		if err := client.Connect(w, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		client.Run()
	}
}
