package actions

import (
	"fmt"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Tailwind = MakeTailwind()

func MakeTailwind() game.Action {
	config := game.ActionConfig{
		Name:  "Tailwind",
		Jutsu: game.Ninjutsu,
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

				for _, tx := range g.Modifiers {
					if tx.Context.SourcePlayerID == nil {
						continue
					}

					if *tx.Context.SourcePlayerID == *context.SourcePlayerID && tx.Mutation.ID == modifiers.SpeedUpTeam.ID {
						g.PushLog(fmt.Sprintf("%s failed.", config.Name))
						return transactions
					}

				}

				su := modifiers.SpeedUpTeam
				su.Duration = 5
				modifiers := []game.Modifier{su}
				mutation := mutations.AddModifiers(modifiers...)
				context.ParentActorID = nil
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
