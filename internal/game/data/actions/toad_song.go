package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var ToadSong = MakeToadSong()

func MakeToadSong() game.Action {
	nature := game.NsYang
	config := game.ActionConfig{
		Name:        "Toad Song",
		Nature:      &nature,
		Jutsu:       game.Senjutsu,
		Description: "Inverts the speed of all active shinobi.",
	}
	return game.Action{
		ID:              uuid.MustParse("02796a9b-add5-5a5c-a01b-5bc6e26d0135"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPrioritySlow,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				for _, tx := range g.Modifiers {
					if tx.Context.SourcePlayerID == nil {
						continue
					}

					if *tx.Context.SourcePlayerID == *context.SourcePlayerID && tx.Mutation.ID == modifiers.ToadSong.ID {
						log := game.NewLogContext("$action$ failed.", context)
						log_tx := game.MakeTransaction(game.AddLogs(log), context)
						return append(transactions, log_tx)
					}
				}

				mod := modifiers.ToadSong
				mod.Duration = 5
				mutation := mutations.AddModifiers(false, mod)
				context.ParentActorID = nil
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
