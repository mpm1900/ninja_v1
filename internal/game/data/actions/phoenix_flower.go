package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var PhoenixFlower = MakePhoenixFlower()

func MakePhoenixFlower() game.Action {
	config := game.ActionConfig{
		Name:        "Phoenix Flower",
		Description: "Hits up-to 6 times. High critical chance.",
		Accuracy:    game.Ptr(85),
		Power:       game.Ptr(20),
		Stat:        game.Ptr(game.StatAttack),
		Nature:      game.Ptr(game.NsFire),
		Cost:        game.Ptr(0),
		TargetCount: game.Ptr(1),
		Jutsu:       game.Bukijutsu,
		CritChance:  game.Ptr(15),
		CritMod:     1.5,
	}
	return game.Action{
		ID:              uuid.MustParse("4ac4894c-2ff3-5142-b087-a8924837cefc"),
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

				conf, _ := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damage_config := game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor())
				damage_config.Repeat = true
				damage_config.RepeatMax = 6
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
