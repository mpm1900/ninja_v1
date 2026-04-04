package instance

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

type RequestType = string

const (
	AddActor         RequestType = "add-actor"
	RemoveActor      RequestType = "remove-actor" //TEMP
	PushAction       RequestType = "push-action"
	RemoveAction     RequestType = "remove-action"
	SetActorPlayer   RequestType = "set-actor-player"   // TEMP
	SetActorPosition RequestType = "set-actor-position" // TEMP
	RunGameActions   RequestType = "run-game-actions"   // TEMP
	ValidateState    RequestType = "validate-state"     // TEMP
	ResolvePrompt    RequestType = "resolve-prompt"

	GetTargets      RequestType = "get-targets"
	ValidateContext RequestType = "validate-context"
)

type Request struct {
	Type          RequestType  `json:"type"`
	ClientID      uuid.UUID    `json:"client_ID"`
	PromptID      *uuid.UUID   `json:"prompt_ID"`
	ModifierID    *uuid.UUID   `json:"modifier_ID"`
	PositionIndex *int         `json:"position_index"`
	Context       game.Context `json:"context"`
}
