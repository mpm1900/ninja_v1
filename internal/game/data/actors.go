package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actors"

	"github.com/google/uuid"
)

var ACTORS map[uuid.UUID]game.ActorDef = map[uuid.UUID]game.ActorDef{
	actors.Itachi.ActorID:     actors.Itachi,
	actors.Sasuke.ActorID:     actors.Sasuke,
	actors.Kisame.ActorID:     actors.Kisame,
	actors.Kakuzu.ActorID:     actors.Kakuzu,
	actors.Hidan.ActorID:      actors.Hidan,
	actors.Naruto.ActorID:     actors.Naruto,
	actors.Minato.ActorID:     actors.Minato,
	actors.Jiraiya.ActorID:    actors.Jiraiya,
	actors.Orochimaru.ActorID: actors.Orochimaru,
	actors.Yamato.ActorID:     actors.Yamato,
	actors.Raikage.ActorID:    actors.Raikage,
	actors.Kaguya.ActorID:     actors.Kaguya,
	actors.Guy.ActorID:        actors.Guy,
	actors.Kakashi.ActorID:    actors.Kakashi,
}
