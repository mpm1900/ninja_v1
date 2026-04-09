package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var DragonDance = MakeDragonDance()

func MakeDragonDance() game.Action {
	nature := game.NsTai
	config := game.ActionConfig{
		Name:        "Dragon Dance",
		Nature:      &nature,
		Jutsu:       game.Taijutsu,
		Description: "Raises the user's Speed and Physical Attack stats.",
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}
				tju := modifiers.AttackUpSource
				su := modifiers.SpeedUpSource

				modifiers := []game.Modifier{
					tju,
					su,
				}
				mutation := mutations.AddModifiers(false, false, modifiers...)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
