package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var PowerBoost = MakePowerBoost()

func MakePowerBoost() game.Action {
	nature := game.NsPure
	config := game.ActionConfig{
		Name:        "Power Boost",
		Nature:      &nature,
		Jutsu:       game.Ninjutsu,
		Description: "Powers up target's attacks this turn.",
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.ComposeAF(game.TeamFilter, game.OtherFilter, game.ActiveFilter, game.AliveFilter),
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityP1,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				mutation := mutations.BoostActionPower(1.5)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
