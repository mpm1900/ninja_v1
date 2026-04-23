package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Expansion = MakeExpansion()

func MakeExpansion() game.Action {
	config := game.ActionConfig{
		Name:        "Expansion",
		Nature:      game.Ptr(game.NsYang),
		Jutsu:       game.Taijutsu,
		Description: "Raises the user's Attack and Defense stats.",
	}
	return game.Action{
		ID:              uuid.MustParse("435490c1-ede2-5875-9edf-1c36d4917741"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}
				mutation := mutations.AddModifiers(false, modifiers.AttackUpSource, modifiers.DefenseUpSource)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
