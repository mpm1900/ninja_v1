package actions

import (
	"math/rand/v2"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var KebariSenbon = MakeKebariSenbon()

func MakeKebariSenbon() game.Action {
	config := game.ActionConfig{
		Name:        "Kebari Senbon",
		Description: "Hits 2-6 times. High critical chance.",
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(20),
		Stat:        game.Ptr(game.StatChakraAttack),
		Nature:      game.Ptr(game.NsYang),
		Cost:        game.Ptr(50),
		TargetCount: game.Ptr(1),
		Jutsu:       game.Bukijutsu,
		CritChance:  game.Ptr(getCriticalStage(1)),
		CritMod:     1.5,
	}
	return game.Action{
		ID:              uuid.MustParse("0de3affc-7513-41b0-8622-c603ccb8ee8a"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            modifiers.UseStaminaCost(*config.Cost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf, _ := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damage_config := game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor())
				damage_config.Repeat = true
				damage_config.RepeatMax = rand.IntN(4) + 2
				damages := game.NewDamage(conf, damage_config)
				transactions = append(
					transactions,
					game.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
