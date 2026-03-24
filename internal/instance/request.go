package instance

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

type RequestType = string

const (
	AddActor       RequestType = "add-actor"
	RemoveActor    RequestType = "remove-actor"
	AddModifier    RequestType = "add-modifier"
	RemoveModifier RequestType = "remove-modifier"
	PushAction     RequestType = "push-action"
)

type Request struct {
	Type       RequestType  `json:"type"`
	ClientID   uuid.UUID    `json:"client_ID"`
	ActionID   *uuid.UUID   `json:"action_ID"`
	ModifierID *uuid.UUID   `json:"modifier_ID"`
	Context    game.Context `json:"context"`
}
