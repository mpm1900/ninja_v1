package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var MODIFIERS = map[uuid.UUID]game.Modifier{
	modifiers.Protected.ID:      modifiers.Protected,
	modifiers.Stunned.ID:        modifiers.Stunned,
	modifiers.Rage.ID:           modifiers.Rage,
	modifiers.AttackUpSource.ID: modifiers.AttackUpSource,
	modifiers.JutsuUpSource.ID:  modifiers.JutsuUpSource,
	modifiers.SpeedUpSource.ID:  modifiers.SpeedUpSource,
	modifiers.SpeedUpAll.ID:     modifiers.SpeedUpAll,
	modifiers.WaterAbsorb.ID:    modifiers.WaterAbsorb,
}
