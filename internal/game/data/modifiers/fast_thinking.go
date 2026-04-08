package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var fastThinkingID = uuid.New()
var FastThinking = game.Modifier{
	ID:       uuid.New(),
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
