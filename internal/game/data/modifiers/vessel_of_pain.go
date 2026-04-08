package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var vesselOfPainID = uuid.New()

var VesselOfPain game.Modifier = game.Modifier{
	ID:       vesselOfPainID,
	GroupID:  &vesselOfPainID,
	Name:     "Vessel of Pain",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&vesselOfPainID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.OtherFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.Stats[game.StatChakraAttack] = game.Round(float64(actor.Stats[game.StatChakraAttack]) * 0.75)
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
