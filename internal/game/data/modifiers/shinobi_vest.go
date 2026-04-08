package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var shinobiVestID = uuid.New()

var ShinobiVest game.Modifier = game.Modifier{
	ID:       shinobiVestID,
	GroupID:  &shinobiVestID,
	Name:     "Shinobi Vest",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&shinobiVestID,
			game.MutPriorityPostStagedStats,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				for i := range actor.Actions {
					actor.Actions[i].Disabled = actor.Actions[i].ID != game.Switch.ID && actor.Actions[i].Config.Power == nil
				}

				actor.Stats[game.StatChakraDefense] = game.Round(float64(actor.Stats[game.StatChakraDefense]) * 1.5)
				return actor
			},
		),
	},
	Triggers: []game.Trigger{},
}
