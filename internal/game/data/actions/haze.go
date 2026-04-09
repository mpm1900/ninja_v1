package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Haze = MakeHaze()

func MakeHaze() game.Action {
	nature := game.NsIce
	config := game.ActionConfig{
		Name:        "Haze",
		Nature:      &nature,
		Jutsu:       game.Ninjutsu,
		Description: "Nullified all stat ups/downs.",
	}
	return game.Action{
		ID:              uuid.MustParse("63db7718-b73b-5f31-8b1f-c2dfa5bd5c65"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				for _, tx := range g.Modifiers {
					if tx.Context.SourcePlayerID == nil {
						continue
					}

					if *tx.Context.SourcePlayerID == *context.SourcePlayerID && tx.Mutation.ID == modifiers.Haze.ID {
						log := game.NewLogContext("$action$ failed.", context)
						log_tx := game.MakeTransaction(game.AddLogs(log), context)
						return append(transactions, log_tx)
					}

				}

				su := modifiers.Haze
				su.Duration = 5
				modifiers := []game.Modifier{su}
				mutation := mutations.AddModifiers(false, false, modifiers...)
				context.ParentActorID = nil
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
