package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var tauntedID = uuid.New()
var Taunted = game.Modifier{
	ID:       uuid.New(),
	Name:     "Taunted",
	Icon:     "taunted",
	GroupID:  &tauntedID,
	Duration: 4,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&tauntedID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.TargetFilter),
			func(g game.Game, a game.Actor, c game.Context) game.Actor {
				for i := range a.Actions {
					a.Actions[i].Disabled = a.Actions[i].ID != game.Switch.ID && a.Actions[i].Config.Power == nil
				}

				return a
			},
		),
	},
}
