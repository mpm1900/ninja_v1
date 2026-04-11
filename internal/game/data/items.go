package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var ITEMS map[uuid.UUID]game.Modifier = map[uuid.UUID]game.Modifier{
	modifiers.SealOfBodyProtection.ID: modifiers.SealOfBodyProtection,
	modifiers.SealOfImmortality.ID:    modifiers.SealOfImmortality,
	modifiers.CurseMarkOfChakra.ID:    modifiers.CurseMarkOfChakra,
	modifiers.CurseMarkOfSpeed.ID:     modifiers.CurseMarkOfSpeed,
	modifiers.CurseMarkOfStrength.ID:  modifiers.CurseMarkOfStrength,
	modifiers.FirstAid.ID:             modifiers.FirstAid,
	modifiers.GedoShard.ID:            modifiers.GedoShard,
	modifiers.Onigiri.ID:              modifiers.Onigiri,
	modifiers.ShinobiVest.ID:          modifiers.ShinobiVest,
}
