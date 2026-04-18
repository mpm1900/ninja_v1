package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Rasengan = MakeRasengan()
var RasenganRecharge = MakeRasenganRecharge()

func MakeRasengan() game.Action {
	config := game.ActionConfig{
		Name:        "Rasengan",
		Description: "Powerful chakra attack. Must recharge the next turn.",
		Nature:      game.Ptr(game.NsPure),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(150),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(100),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}
	return game.Action{
		ID:              uuid.MustParse("e0874a45-2f62-5544-a4a2-f440644407db"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(*config.Cost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, 1))
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
		Name:        "Recharging...",
		LogSuccessF: &logf,
	}
	return game.Action{
		ID:              uuid.MustParse("2eaa6398-06a5-56fe-b90d-e9db6f044744"),
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
