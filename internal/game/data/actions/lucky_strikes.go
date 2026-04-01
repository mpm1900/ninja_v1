package actions

import (
	"fmt"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var LuckyStrikes = MakeLuckyStrikes()

func MakeLuckyStrikes() game.Action {
	accuracy := 60
	power := 20
	stat := game.Attack
	nature := game.NsTai
	chakraCost := 30
	config := game.ActionConfig{
		Name:     "Lucky Strikes",
		Accuracy: &accuracy,
		Power:    &power,
		Stat:     &stat,
		Nature:   &nature,
		Cost:     &chakraCost,
		Jutsu:    game.Taijutsu,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter, game.AliveFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		Cost:            mutations.UseStaminaSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.AllGameFilter,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				accuracy := *config.Accuracy
				for {
					config.Accuracy = nil
					roll := game.MakeActionRoll()
					if roll > accuracy {
						zero := 0
						config.Accuracy = &zero
						fmt.Println("ROLL", roll, accuracy)
					} else {
						lx := game.MakeTransaction(game.AddLogs("It hits!"), context)
						transactions = append(transactions, lx)
					}

					damages := mutations.NewDamage(config, game.NewDamageConfig())
					transactions = append(
						transactions,
						mutations.MakeDamageTransactions(context, damages)...,
					)

					if config.Accuracy != nil && *config.Accuracy == 0 {
						break
					}
				}

				return transactions
			},
		},
	}
}
