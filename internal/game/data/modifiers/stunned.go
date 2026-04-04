package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var stunnedID = uuid.New()
var Stunned = game.Modifier{
	ID:       uuid.New(),
	Name:     "Stunned",
	GroupID:  &stunnedID,
	Duration: 0,
	Mutations: []game.ModifierMutation{
		game.MakeActorMutation(
			&stunnedID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.TargetFilter),
			func(g game.Game, a game.Actor, c game.Context) game.Actor {
				a.Stunned = true
				return a
			},
		),
	},
}
