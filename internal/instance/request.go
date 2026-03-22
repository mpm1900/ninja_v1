package instance

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

type Request struct {
	Type       string       `json:"type"`
	ClientID   uuid.UUID    `json:"clientID"`
	ModifierID *uuid.UUID   `json:"modifierID"`
	Context    game.Context `json:"context"`
}
