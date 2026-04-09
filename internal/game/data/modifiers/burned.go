package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var burnedID = uuid.MustParse("497240d3-0d1a-5194-b475-62e418ad94eb")

var BurnedTrigger game.Trigger = game.Trigger{
	ID:         uuid.MustParse("6d3a4475-cca7-57e1-b606-3c4515c955ec"),
	ModifierID: burnedID,
	On:         game.OnTurnEnd,
	Check: func(p, g game.Game, context game.Context, tx game.Transaction[game.Modifier]) bool {
		return true
	},
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			context.TargetPositionIDs = []uuid.UUID{}
			context.TargetActorIDs = []uuid.UUID{*context.SourceActorID}
			mut := mutations.RatioDamage(0.125)
			return []game.Transaction[game.GameMutation]{
				game.MakeTransaction(mut, context),
			}
		},
	},
}

var Burned game.Modifier = game.Modifier{
	ID:       burnedID,
	GroupID:  &burnedID,
	Name:     "Burned",
	Icon:     "burned",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&burnedID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				if actor.Statused {
					return actor
				}

				actor.Statused = true
				actor.Stats[game.StatAttack] = game.Round(float64(actor.Stats[game.StatAttack]) * 0.5)
				return actor
			},
		),
	},
	Triggers: []game.Trigger{
		BurnedTrigger,
	},
}
