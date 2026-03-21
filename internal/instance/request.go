package instance

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

type Request struct {
	Type     string
	ClientID uuid.UUID
	Context  game.Context
}
