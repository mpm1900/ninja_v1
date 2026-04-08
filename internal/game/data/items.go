package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var ITEMS map[uuid.UUID]game.Modifier = map[uuid.UUID]game.Modifier{
	modifiers.BodyProtectionSeal.ID:  modifiers.BodyProtectionSeal,
	modifiers.CurseMarkOfChakra.ID:   modifiers.CurseMarkOfChakra,
	modifiers.CurseMarkOfSpeed.ID:    modifiers.CurseMarkOfSpeed,
	modifiers.CurseMarkOfStrength.ID: modifiers.CurseMarkOfStrength,
	modifiers.Leftovers.ID:           modifiers.Leftovers,
	modifiers.LifeOrb.ID:             modifiers.LifeOrb,
	modifiers.Onigiri.ID:             modifiers.Onigiri,
	modifiers.ShinobiVest.ID:         modifiers.ShinobiVest,
}
