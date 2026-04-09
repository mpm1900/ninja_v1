package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var firstAidID = uuid.New()

var FirstAidTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: firstAidID,
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
			mut := mutations.RatioHeal(0.0625)
			return []game.Transaction[game.GameMutation]{
				game.MakeTransaction(mut, context),
			}
		},
	},
}

var FirstAid game.Modifier = game.Modifier{
	ID:       firstAidID,
	GroupID:  &firstAidID,
	Name:     "First Aid Kit",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&firstAidID),
	},
	Triggers: []game.Trigger{
		FirstAidTrigger,
	},
}
