package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var feelingOfPainID = uuid.MustParse("a1adcca5-e175-4454-a615-3e8be849e8ab")

var FeelingOfPain game.Modifier = game.Modifier{
	ID:                feelingOfPainID,
	GroupID:           &feelingOfPainID,
	Icon:              "std_strength",
	Name:              "Feeling of Pain",
	Description:       "Attack x0.75.",
	ParentDescription: "Other shinobi: Attack x0.75.",
	Show:              true,
	Duration:          game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&feelingOfPainID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.OtherFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.Stats[game.StatAttack] = game.Round(float64(actor.Stats[game.StatAttack]) * 0.75)
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
