package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var StunnedID = uuid.New()
var Stunned = game.Modifier{
	ID:       uuid.New(),
	Name:     "Stunned",
	GroupID:  StunnedID,
	Duration: 0,
	Mutations: []game.ActorMutation{
		game.MakeActorMutation(
			&StunnedID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.TargetFilter),
			func(a game.Actor, c game.Context) game.Actor {
				a.Stunned = true
				return a
			},
		),
	},
}
