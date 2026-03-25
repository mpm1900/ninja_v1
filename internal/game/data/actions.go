package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)


var ACTIONS map[uuid.UUID]game.Action = map[uuid.UUID]game.Action{
	actions.LeafJab.ID:     actions.LeafJab,
	actions.DragonDance.ID: actions.DragonDance,
	actions.Fireball.ID:    actions.Fireball,
	actions.Chidori.ID:     actions.Chidori,
}
