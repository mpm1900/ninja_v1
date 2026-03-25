package instance

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

type RequestType = string

const (
	AddActor         RequestType = "add-actor"
	RemoveActor      RequestType = "remove-actor"
	AddModifier      RequestType = "add-modifier"    // TEMP
	RemoveModifier   RequestType = "remove-modifier" // TEMP
	PushAction       RequestType = "push-action"
	SetActorPlayer   RequestType = "set-actor-player"   // TEMP
	SetActorPosition RequestType = "set-actor-position" // TEMP
)

type Request struct {
	Type          RequestType  `json:"type"`
	ClientID      uuid.UUID    `json:"client_ID"`
	ActionID      *uuid.UUID   `json:"action_ID"`
	ModifierID    *uuid.UUID   `json:"modifier_ID"`
	PositionIndex *int         `json:"position_index"`
	Context       game.Context `json:"context"`
}
