package data

import (
	"ninja_v1/internal/game"
	modifiers "ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var MODIFIERS = map[uuid.UUID]game.Modifier{
	modifiers.GenjutsuUpSource.ID: modifiers.GenjutsuUpSource,
	modifiers.SpeedUpSource.ID:    modifiers.SpeedUpSource,
	modifiers.SpeedUpAll.ID:       modifiers.SpeedUpAll,
	modifiers.TaijutsuUpSource.ID: modifiers.TaijutsuUpSource,
}
