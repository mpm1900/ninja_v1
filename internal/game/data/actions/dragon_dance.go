package actions

import (
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var DragonDance = MakeDragonDance()

func MakeDragonDance() game.Action {
	accuracy := 100
	config := game.ActionConfig{
		Name:     "Dragon Dance",
		Accuracy: &accuracy,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: 0,
			Filter:   game.AllGameFilter,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				modifiers := []game.Modifier{
					modifiers.TaijutsuUpSource,
					modifiers.SpeedUpSource,
				}
				mutation := mutations.AddModifiers(modifiers)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
