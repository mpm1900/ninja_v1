package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actors"

	"github.com/google/uuid"
)

var ACTORS map[uuid.UUID]game.ActorDef = map[uuid.UUID]game.ActorDef{
	actors.Madara.ActorID:     actors.Madara,
	actors.Itachi.ActorID:     actors.Itachi,
	actors.Sasuke.ActorID:     actors.Sasuke,
	actors.Shisui.ActorID:     actors.Shisui,
	actors.Kisame.ActorID:     actors.Kisame,
	actors.Kakuzu.ActorID:     actors.Kakuzu,
	actors.Hidan.ActorID:      actors.Hidan,
	actors.Deidara.ActorID:    actors.Deidara,
	actors.Hashirama.ActorID:  actors.Hashirama,
	actors.Naruto.ActorID:     actors.Naruto,
	actors.Minato.ActorID:     actors.Minato,
	actors.Pain.ActorID:       actors.Pain,
	actors.Jiraiya.ActorID:    actors.Jiraiya,
	actors.Orochimaru.ActorID: actors.Orochimaru,
	actors.Yamato.ActorID:     actors.Yamato,
	actors.Raikage.ActorID:    actors.Raikage,
	actors.Guy.ActorID:        actors.Guy,
	actors.RockLee.ActorID:    actors.RockLee,
	actors.Kakashi.ActorID:    actors.Kakashi,
}
