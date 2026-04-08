package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var lifeOrbID = uuid.New()

var LifeOrbTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: lifeOrbID,
	On:         game.OnTurnEnd,
	Check: func(p, g game.Game, context game.Context, tx game.Transaction[game.Modifier]) bool {
		return true
	},
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			context.TargetPositionIDs = []uuid.UUID{}
			context.TargetActorIDs = []uuid.UUID{*context.SourceActorID}
			mut := mutations.RatioDamage(0.1)
			return []game.Transaction[game.GameMutation]{
				game.MakeTransaction(mut, context),
			}
		},
	},
}

var LifeOrb game.Modifier = game.Modifier{
	ID:       lifeOrbID,
	GroupID:  &lifeOrbID,
	Name:     "Life Orb",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&lifeOrbID,
			game.MutPriorityDefault,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				if actor.Statused {
					return actor
				}

				actor.DamageMult += 0.3
				return actor
			},
		),
	},
	Triggers: []game.Trigger{
		LifeOrbTrigger,
	},
}
