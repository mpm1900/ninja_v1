package data

import (
	"ninja_v1/internal/game"
	modifiers "ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var GenjustuUpSource = modifiers.NewStatDoublerSource(game.StatGenjutsu, "Genjustu Up")

var MODIFIERS = map[uuid.UUID]game.Modifier{
	GenjustuUpSource.ID: GenjustuUpSource,
}
