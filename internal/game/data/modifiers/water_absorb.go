package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var waterAbsorbID = uuid.New()

var WaterAbsorb game.Modifier = game.Modifier{
	ID:       waterAbsorbID,
	Name:     "Water Absorb",
	GroupID:  waterAbsorbID,
	Duration: game.ModifierDurationInf,
	Mutations: []game.ActorMutation{
		game.MakeActorMutation(
			&waterAbsorbID,
			game.MutPrioritySet,
			game.ComposeAF(game.SourceFilter),
			func(actor game.Actor, context game.Context) game.Actor {
				actor.NatureResistance[game.NatureWater] = -0.5
				return actor
			},
		),
	},
}
