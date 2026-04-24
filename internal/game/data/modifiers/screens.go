package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var pddtID = uuid.New()
var PhysicalDamageDownTeam = game.Modifier{
	ID:          uuid.New(),
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
