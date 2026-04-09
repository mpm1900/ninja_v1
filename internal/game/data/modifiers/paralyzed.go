package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var paralysisID = uuid.MustParse("b269dcbc-a6a0-5ea3-9ed5-6e8b5fee8024")

var Paralysis game.Modifier = game.Modifier{
	ID:       paralysisID,
	GroupID:  &paralysisID,
	Name:     "Paralyzed",
	Icon:     "paralyzed",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&paralysisID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				if actor.Statused {
					return actor
				}

				actor.Statused = true
				actor.Paralyzed = true
				actor.Stats[game.StatSpeed] = game.Round(float64(actor.Stats[game.StatSpeed]) * 0.25)
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
