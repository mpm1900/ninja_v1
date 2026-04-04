package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var waterAbsorbID = uuid.New()

var WaterAbsorb game.Modifier = game.Modifier{
	ID:       waterAbsorbID,
	Name:     "Water Absorb",
	GroupID:  &waterAbsorbID,
	Duration: game.ModifierDurationInf,
	Mutations: []game.ModifierMutation{
		{
			ModifierGroupID: &waterAbsorbID,
			ActorFilter: func(g game.Game, a game.Actor, c game.Context) bool {
				return game.SourceFilter(g, a, c)
			},
			ActorDelta: func(g game.Game, a game.Actor, c game.Context) game.Actor {
				action, ok := g.GetActiveAction()
				if !ok || action.Config.Nature == nil {
					return a
				}

				if *action.Config.Nature != game.NsWater {
					return a
				}

				a.NatureResistance[game.NatureWater] = -1.0
				return a
			},
			GameMutation: game.GameMutation{
				Priority: game.MutPrioritySet,
				Delta: func(g game.Game, context game.Context) game.Game {
					return g
				},
			},
		},
	},
}
