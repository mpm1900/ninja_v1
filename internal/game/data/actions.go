package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var LeafJab = actions.MakeLeafJab()
var DragonDance = actions.MakeDragonDance()
var Fireball = actions.MakeFireball()

var ACTIONS map[uuid.UUID]game.Action = map[uuid.UUID]game.Action{
	LeafJab.ID:     LeafJab,
	DragonDance.ID: DragonDance,
	Fireball.ID:    Fireball,
}
