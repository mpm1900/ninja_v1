package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Amaterasu = MakeAmaterasu()

func MakeAmaterasu() game.Action {
	ID := uuid.MustParse("d103e605-9381-52fd-9cb8-450b7315a9f9")
	nature := game.NsYin

	config := game.ActionConfig{
		Name:        "Amaterasu",
		Description: "Burns taraget.",
		Nature:      &nature,
		Stat:        game.Ptr(game.ChakraAttack),
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
				damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, 1))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				targets := g.GetTargets(context)
				for _, target := range targets {
					mut_ctx := game.Context{
						SourcePlayerID: &target.PlayerID,
						SourceActorID:  &target.ID,
						ParentActorID:  nil, // do not remove on switch
						TargetActorIDs: []uuid.UUID{target.ID},
					}
					mutation := mutations.AddModifiers(true, modifiers.Burned)
					transaction := game.MakeTransaction(mutation, mut_ctx)
					transactions = append(transactions, transaction)
				}

				return transactions
			},
		},
	}
}
