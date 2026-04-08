package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var speedBoostID = uuid.New()
var SpeedBoostTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: speedBoostID,
	On:         game.OnTurnEnd,
	Check: func(g1, g2 game.Game, ctx game.Context, t game.Transaction[game.Modifier]) bool {
		return true
	},
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			transactions := []game.GameTransaction{}

			mutation := mutations.AddModifiers(false, SpeedUpSource)
			transactions = append(transactions, game.MakeTransaction(mutation, context))
			return transactions
		},
	},
}

var SpeedBoost game.Modifier = game.Modifier{
	ID:       speedBoostID,
	GroupID:  &speedBoostID,
	Name:     "Speed Boost",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&speedBoostID),
	},
	Triggers: []game.Trigger{
		SpeedBoostTrigger,
	},
}
