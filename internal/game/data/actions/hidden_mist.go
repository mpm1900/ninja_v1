package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var HiddenMist = MakeHiddenMist()

func MakeHiddenMist() game.Action {
	config := game.ActionConfig{
		Name:        "Hidden Mist",
		Nature:      game.Ptr(game.NsWater),
		Jutsu:       game.Ninjutsu,
		Description: "All non-water shinobi have Accuracy down x2.",
	}
	return game.Action{
		ID:              uuid.MustParse("a8c7fab6-c3e0-4933-ab1a-3d05376c05b1"),
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

					if *tx.Context.SourcePlayerID == *context.SourcePlayerID && tx.Mutation.ID == modifiers.HiddenMist.ID {
						log := game.NewLogContext("$action$ failed.", context)
						log_tx := game.MakeTransaction(game.AddLogs(log), context)
						return append(transactions, log_tx)
					}

				}

				su := modifiers.HiddenMist
				su.Duration = 5
				modifiers := []game.Modifier{su}
				mutation := mutations.AddModifiers(false, modifiers...)
				context.ParentActorID = nil
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
