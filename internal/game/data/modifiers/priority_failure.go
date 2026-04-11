package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var priorityFailureID = uuid.MustParse("7fdf20ca-b003-5082-8980-a0c6990169d0")
var PriorityFailure = game.Modifier{
	ID:          uuid.MustParse("20070a28-dc4d-5aa8-a45c-27d952d221a6"),
	Name:        "Priority Failure",
	Description: "Enemy non-attacking priority actions fail.",
	Show:        true,
	GroupID:     &priorityFailureID,
	Duration:    0,
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
