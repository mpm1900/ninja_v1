package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var protectedID = uuid.New()
var Protected = game.Modifier{
	ID:       uuid.New(),
	Name:     "Protected",
	GroupID:  &protectedID,
	Duration: 0,
	Mutations: []game.ModifierMutation{
		game.MakeActorMutation(
			&protectedID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.SourceFilter),
			func(g game.Game, a game.Actor, c game.Context) game.Actor {
				a.Protected = true
				return a
			},
		),
	},
}
