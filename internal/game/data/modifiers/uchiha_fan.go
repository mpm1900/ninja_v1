package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var uchihaFanID = uuid.MustParse("d0676564-df6a-4716-a3c8-7ac4c5aaabff")
var UchihaFan game.Modifier = game.Modifier{
	ID:          uchihaFanID,
	GroupID:     &uchihaFanID,
	Icon:        "uchiha_fan",
	Name:        "Uchiha Fan",
	Description: "Fire attacks deal 10% more damage.",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&uchihaFanID,
			game.MutPriorityDefault,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.NatureDamage[game.NatureFire] += 0.1
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
