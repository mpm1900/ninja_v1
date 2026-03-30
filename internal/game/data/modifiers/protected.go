package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var protectedCooldown = 1
var ProtectedID = uuid.New()
var Protected = game.Modifier{
	ID:       uuid.New(),
	Name:     "Protected",
	GroupID:  ProtectedID,
	Duration: &protectedCooldown,
	Mutations: []game.ActorMutation{
		game.MakeActorMutation(
			&ProtectedID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.SourceFilter),
			func(a game.Actor, c game.Context) game.Actor {
				a.Protected = true
				return a
			},
		),
	},
}
