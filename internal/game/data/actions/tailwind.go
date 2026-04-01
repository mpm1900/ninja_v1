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
	nature := game.NsWind
	config := game.ActionConfig{
		Name:        "Tailwind",
		Nature:      &nature,
		Jutsu:       game.Ninjutsu,
		Description: "Doubles the speed of the user's party for 4 turns. This effect cannot stack.",
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
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				for _, tx := range g.Modifiers {
					if tx.Context.SourcePlayerID == nil {
						continue
					}

					if *tx.Context.SourcePlayerID == *context.SourcePlayerID && tx.Mutation.ID == modifiers.Tailwind.ID {
						log_tx := game.MakeTransaction(game.AddLogs(fmt.Sprintf("%s failed.", config.Name)), context)
						return append(transactions, log_tx)
					}

				}

				su := modifiers.Tailwind
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
