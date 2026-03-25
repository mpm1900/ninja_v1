package mutations

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewNoop(groupID *uuid.UUID) game.ModifierMutation {
	return game.MakeModifierMutation(
		groupID,
		0,
		game.SourceFilter,
		func(a game.Actor, c game.Context) game.Actor {
			return a
		},
	)
}
