package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Graft = MakeGraft()

func MakeGraft() game.Action {
	ID := uuid.MustParse("fdbfd320-071a-46a2-b449-e1455d1a3d14")

	config := game.ActionConfig{
		Name:        "Graft",
		Description: "Heals an ally or damages an enemy.",
		Nature:      game.Ptr(game.NsYang),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(70),
		Stat:        game.Ptr(game.ChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(1),
		Jutsu:       game.Ninjutsu,
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
				game.SourceIsActionOffCooldown,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)

				for _, target := range g.GetTargets(context) {
					isTeam := context.SourcePlayerID != nil && target.PlayerID == *context.SourcePlayerID
					ctx := context
					ctx.TargetPositionIDs = []uuid.UUID{*target.PositionID}
					if isTeam {
						damages := mutations.RatioHeal(0.5)
						transactions = append(
							transactions,
							mutations.MakeDamageTransactions(ctx, damages)...,
						)
					} else {
						damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, 1))
						transactions = append(
							transactions,
							mutations.MakeDamageTransactions(ctx, damages)...,
						)
					}
				}

				return transactions
			},
		},
	}
}
