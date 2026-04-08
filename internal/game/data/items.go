package data

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var ITEMS map[uuid.UUID]game.Modifier = map[uuid.UUID]game.Modifier{}
