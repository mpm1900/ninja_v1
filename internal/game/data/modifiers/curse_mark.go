package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var cmoChakraID = uuid.New()
var cmoSpeedID = uuid.New()
var cmoStrengthID = uuid.New()

var CurseMarkOfChakra game.Modifier = game.Modifier{
	ID:       cmoChakraID,
	GroupID:  &cmoChakraID,
	Name:     "Curse Mark of Chakra",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&cmoChakraID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.ActionLocked = true
				actor.Stats[game.StatChakraAttack] = game.Round(float64(actor.Stats[game.StatChakraAttack]) * 1.5)
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}

var CurseMarkOfSpeed game.Modifier = game.Modifier{
	ID:       cmoSpeedID,
	GroupID:  &cmoSpeedID,
	Name:     "Curse Mark of Speed",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&cmoSpeedID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.ActionLocked = true
				actor.Stats[game.StatSpeed] = game.Round(float64(actor.Stats[game.StatSpeed]) * 1.5)
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}

var CurseMarkOfStrength game.Modifier = game.Modifier{
	ID:       cmoStrengthID,
	GroupID:  &cmoStrengthID,
	Name:     "Curse Mark of Strength",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&cmoStrengthID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.ActionLocked = true
				actor.Stats[game.StatAttack] = game.Round(float64(actor.Stats[game.StatAttack]) * 1.5)
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
