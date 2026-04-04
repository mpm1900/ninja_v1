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
			ActorFilter: func(a game.Actor, c game.Context) bool {
				return game.SourceFilter(a, c)
			},
			ActorDelta: func(g game.Game, a game.Actor, c game.Context) game.Actor {
				if g.ActiveContext == nil || g.ActiveContext.ActionID == nil {
					return a
				}

				source, ok := g.GetSource(*g.ActiveContext)
				if !ok {
					return a
				}

				var action *game.Action
				for i := range source.Actions {
					if source.Actions[i].ID == *g.ActiveContext.ActionID {
						action = &source.Actions[i]
						break
					}
				}
				if action == nil || action.Config.Nature == nil {
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
