package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var addImmunitiesID = uuid.New()

func AddImmunities(ids ...uuid.UUID) game.Modifier {
	return game.Modifier{
		ID:      addImmunitiesID,
		GroupID: &addImmunitiesID,
		Name:    "Add Immunity",
		ActorMutations: []game.ActorMutation{
			game.MakeActorMutation(
				&addImmunitiesID,
				game.MutPriorityImmunity,
				game.ComposeAF(game.SourceFilter, game.ActiveFilter),
				func(g game.Game, a game.Actor, c game.Context) game.Actor {
					a.PushImmunities(ids...)
					return a
				},
			),
		},
	}
}
