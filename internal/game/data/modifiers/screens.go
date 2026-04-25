package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var pddtID = uuid.New()
var PhysicalDamageDownTeam = game.Modifier{
	ID:          pddtID,
	Icon:        "physical_reduction_up",
	Name:        "Physical Damage Down",
	Description: "Takes 50% less physical damage.",
	Show:        true,
	GroupID:     &pddtID,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&pddtID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.ActiveFilter, game.TeamFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.DamageReduction[game.Attack] /= 2
				return actor
			},
		),
	},
}

var cddtID = uuid.New()
var ChakraDamageDownTeam = game.Modifier{
	ID:          cddtID,
	Icon:        "chakra_reduction_up",
	Name:        "Chakra Damage Down",
	Description: "Takes 50% less chakra damage.",
	Show:        true,
	GroupID:     &cddtID,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&pddtID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.ActiveFilter, game.TeamFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.DamageReduction[game.ChakraAttack] /= 2
				return actor
			},
		),
	},
}
