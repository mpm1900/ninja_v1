package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var priorityFailureID = uuid.New()
var PriorityFailure = game.Modifier{
	ID:       uuid.New(),
	Name:     "Priority Failure",
	Show:     true,
	GroupID:  &priorityFailureID,
	Duration: 0,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&priorityFailureID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.OtherTeamFilter),
			func(g game.Game, a game.Actor, c game.Context) game.Actor {
				for i, action := range a.Actions {
					if action.Config.Power != nil && action.Priority > 0 {
						a.Actions[i].Filter = game.FalseGameFilter
					}
				}

				return a
			},
		),
	},
}
