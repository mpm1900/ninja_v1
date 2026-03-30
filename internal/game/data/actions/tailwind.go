package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Tailwind = MakeTailwind()

func MakeTailwind() game.Action {
	config := game.ActionConfig{
		Name: "Tailwind",
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

				su := modifiers.SpeedUpTeam
				dur := 5
				su.Duration = &dur
				modifiers := []game.Modifier{su}
				mutation := mutations.AddModifiers(modifiers...)
				suContext := context
				suContext.ParentActorID = nil
				transaction := game.MakeTransaction(mutation, suContext)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
