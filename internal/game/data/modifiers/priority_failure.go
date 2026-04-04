package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var priorityFailureID = uuid.New()
var PriorityFailure = game.Modifier{
	ID:       uuid.New(),
	Name:     "Priority Failure",
	GroupID:  &priorityFailureID,
	Duration: 0,
	Mutations: []game.ModifierMutation{
		game.MakeActorMutation(
			&priorityFailureID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.OtherTeamFilter),
			func(g game.Game, a game.Actor, c game.Context) game.Actor {
				return a
				for _, action := range a.Actions {
					if action.Config.Power != nil && action.Priority > 0 {
						// a.Actions[i].Filter = game.FalseGameFilter
					}
				}

				return a
			},
		),
	},
}
