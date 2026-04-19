package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var SearingMigraine = MakeSearingMigraine()

func MakeSearingMigraine() game.Action {
	ID := uuid.MustParse("dc6edab6-535f-508f-b791-e197283eae86")

	config := game.ActionConfig{
		Name:        "Searing Migraine",
		Description: "Grants the user Fire nature until end of turn.",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(80),
		Stat:        game.Ptr(game.StatChakraAttack),
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
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(*config.Cost),
		MapContext: func(g game.Game, context game.Context) game.Context {
			other_team_actors := g.GetActorsFilters(context, game.ComposeAF(game.ActiveFilter, game.OtherTeamFilter))
			for _, t := range other_team_actors {
				context.TargetPositionIDs = append(context.TargetPositionIDs, *t.PositionID)
			}
			return context
		},
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

				add_mut := mutations.AddModifiers(false, modifiers.AddNature(game.NsFire, 0))
				add_tx := game.MakeTransaction(add_mut, context)
				transactions = append(transactions, add_tx)

				damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor()))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
