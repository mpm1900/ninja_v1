package data

import (
	"ninja_v1/internal/game"
	actors "ninja_v1/internal/game/data/actors"

	"github.com/google/uuid"
)

func GetAllActors() []game.Actor {
	return []game.Actor{
		actors.NewItachi(uuid.Nil, 24),
		actors.NewKisame(uuid.Nil, 24),
		actors.NewKakuzu(uuid.Nil, 24),
	}
}
