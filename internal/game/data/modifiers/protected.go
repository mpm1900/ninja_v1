package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var protectedID = uuid.MustParse("753f21af-6b67-50ff-a386-40528a73e62e")
var Protected = game.Modifier{
	ID:          uuid.MustParse("fa58e0c1-01aa-58ef-b42d-0af89ee4555c"),
	Name:        "Protected",
	Description: "Protected from damage and modifiers.",
	Icon:        "protected",
	Show:        true,
	GroupID:     &protectedID,
	Duration:    0,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&protectedID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.SourceFilter),
			func(g game.Game, a game.Actor, c game.Context) game.Actor {
				a.Protected = true
				a.Safeguarded = true
				return a
			},
		),
	},
}
