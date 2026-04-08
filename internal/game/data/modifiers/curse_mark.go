package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var cmoStrengthID = uuid.New()

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
