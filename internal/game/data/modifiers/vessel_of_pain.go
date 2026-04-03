package modifiers

import (
	"math"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var vesselOfPainID = uuid.New()

var VesselOfPain game.Modifier = game.Modifier{
	ID:       vesselOfPainID,
	GroupID:  &vesselOfPainID,
	Name:     "Vessel of Pain",
	Duration: game.ModifierDurationInf,
	Mutations: []game.ActorMutation{
		game.MakeActorMutation(
			&vesselOfPainID,
			game.MutPriorityDefault,
			game.ComposeAF(game.OtherFilter, game.ActiveFilter),
			func(actor game.Actor, context game.Context) game.Actor {
				actor.Stats[game.StatChakraAttack] = int(math.Floor(float64(actor.Stats[game.StatChakraAttack]) * 0.75))
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
