package actions

import (
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var DragonDance = MakeDragonDance()

func MakeDragonDance() game.Action {
	stat := game.AttackGenjutsu
	config := game.ActionConfig{
		Name: "Dragon Dance",
		Stat: &stat,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.AllGameFilter,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				modifiers := []game.Modifier{
					modifiers.TaijutsuUpSource,
					modifiers.SpeedUpSource,
				}
				mutation := mutations.AddModifiers(modifiers...)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
