package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var fastThinkingID = uuid.MustParse("7a77c7ae-7027-5f3c-9822-7e4406f38891")
var FastThinking = game.Modifier{
	ID:       uuid.MustParse("757200f0-fe4d-5b7d-8de9-1342872a0f2b"),
	Name:     "Fast Thinking",
	Show:     true,
	GroupID:  &fastThinkingID,
	Duration: 0,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&fastThinkingID,
			game.MutPriorityDefault,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter, game.IsAtOrBelowHealthRatio(0.5)),
			func(g game.Game, a game.Actor, c game.Context) game.Actor {
				for i, action := range a.Actions {
					if action.Config.Power == nil && action.Priority == game.ActionPriorityDefault {
						a.Actions[i].Priority = game.ActionPriorityP1
					}
				}

				return a
			},
		),
	},
}
