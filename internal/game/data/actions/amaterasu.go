package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Amaterasu = MakeAmaterasu()

func MakeAmaterasu() game.Action {
	ID := uuid.MustParse("d103e605-9381-52fd-9cb8-450b7315a9f9")

	config := game.ActionConfig{
		Name:        "Amaterasu",
		Description: "Burns target.",
		Nature:      game.Ptr(game.NsYin),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(20),
		Cost:        game.Ptr(30),
		Jutsu:       game.Genjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(*config.Cost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damages := game.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor()))
				transactions = append(
					transactions,
					game.MakeDamageTransactions(context, damages)...,
				)

				targets := g.GetTargets(context)
				for _, target := range targets {
					transactions = append(transactions, applyBurn(target)...)
				}

				return transactions
			},
		},
	}
}
