package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Rasengan = MakeRasengan()
var RasenganRecharge = MakeRasenganRecharge()

func MakeRasengan() game.Action {
	accuracy := 100
	power := 150
	stat := game.ChakraAttack
	nature := game.NsPure
	chakraCost := 100
	config := game.ActionConfig{
		Name:        "Rasengan",
		Description: "Powerful chakra attack. Must recharge the next turn.",
		Accuracy:    &accuracy,
		Power:       &power,
		Stat:        &stat,
		Nature:      &nature,
		Cost:        &chakraCost,
		Jutsu:       game.Ninjutsu,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		Cost:            mutations.UseStaminaSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				damages := mutations.NewDamage(conf, game.NewDamageConfig(1, 1))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				recharge := mutations.QueueAction(RasenganRecharge.ID, context)
				transactions = append(transactions, game.MakeTransaction(recharge, context))

				return transactions
			},
		},
	}
}

func MakeRasenganRecharge() game.Action {
	logf := "%s must recharge."
	config := game.ActionConfig{
		Name:        "Reacharging...",
		LogSuccessF: &logf,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.PositionsLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				return transactions
			},
		},
	}
}
