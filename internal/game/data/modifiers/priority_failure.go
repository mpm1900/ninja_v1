package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var priorityFailureID = uuid.MustParse("7fdf20ca-b003-5082-8980-a0c6990169d0")
var PriorityFailure = game.Modifier{
	ID:          priorityFailureID,
	Name:        "Priority Failure",
	Description: "Enemy priority actions are disabled.",
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
					if !action.Meta.Switch {
						if action.Priority > game.ActionPriorityDefault {
							a.Actions[i].Disabled = true
						}
					}
				}

				return a
			},
		),
	},
}
