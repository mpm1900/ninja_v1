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
	accuracy := 100
	power := 70
	nature := game.NsFire
	stat := game.StatChakraAttack
	targetCount := 0
	chakraCost := 30
	cooldown := 1

	config := game.ActionConfig{
		Name:        "Searing Migraine",
		Description: "Grants the user Fire nature until end of turn.",
		Nature:      &nature,
		Accuracy:    &accuracy,
		Power:       &power,
		Stat:        &stat,
		TargetCount: &targetCount,
		Cost:        &chakraCost,
		Cooldown:    &cooldown,
		Jutsu:       game.Ninjutsu,
	}

	return game.Action{
		ID:              ID,
		Config:          config,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(chakraCost),
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
				add_mut := mutations.AddModifiers(false, modifiers.AddNature(game.NsFire, 0))
				add_tx := game.MakeTransaction(add_mut, context)
				transactions = append(transactions, add_tx)

				damages := mutations.NewDamage(conf, game.NewDamageConfig(1, 1))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
