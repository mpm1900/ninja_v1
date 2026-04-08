package instance

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

type RequestType = string

const (
	AddActor       RequestType = "add-actor"
	RemoveActor    RequestType = "remove-actor" //TEMP
	UpdateActor    RequestType = "update-actor"
	PushAction     RequestType = "push-action"
	RemoveAction   RequestType = "remove-action"
	RunGameActions RequestType = "run-game-actions" // TEMP
	ValidateState  RequestType = "validate-state"   // TEMP
	ResolvePrompt  RequestType = "resolve-prompt"

	GetTargets      RequestType = "get-targets"
	ValidateContext RequestType = "validate-context"
)

type ActorConfig struct {
	AbilityID *uuid.UUID       `json:"ability_ID"`
	ActionIDs []uuid.UUID      `json:"action_IDs"`
	Focus     *game.ActorFocus `json:"focus"`
}

type Request struct {
	Type        RequestType  `json:"type"`
	ClientID    uuid.UUID    `json:"client_ID"`
	PromptID    *uuid.UUID   `json:"prompt_ID"`
	Context     game.Context `json:"context"`
	ActorConfig *ActorConfig `json:"actor_config"`
}
