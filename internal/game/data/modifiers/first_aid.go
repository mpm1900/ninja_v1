package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var firstAidID = uuid.MustParse("4186b062-e930-51cd-b838-12c618b790c0")

var FirstAidTrigger game.Trigger = game.Trigger{
	ID:         uuid.MustParse("7d6c4561-b6e8-5107-8e8b-3d69b18cb93f"),
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
	ID:          firstAidID,
	GroupID:     &firstAidID,
	Name:        "First Aid Kit",
	Description: "End of turn: heal 1/16 HP.",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&firstAidID),
	},
	Triggers: []game.Trigger{
		FirstAidTrigger,
	},
}
