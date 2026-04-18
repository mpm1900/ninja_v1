package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Surf = MakeSurf()

func MakeSurf() game.Action {
	ID := uuid.MustParse("74d5a7d7-cb62-58b4-9ace-e80bf7f0fd40")
	accuracy := 100
	power := 90
	nature := game.NsWater
	stat := game.StatChakraAttack
	targetCount := 0
	cost := 30

	config := game.ActionConfig{
		Name:        "Surf",
		Nature:      &nature,
		Accuracy:    &accuracy,
		Power:       &power,
		Stat:        &stat,
		TargetCount: &targetCount,
		Cost:        &cost,
		Jutsu:       game.Ninjutsu,
	}

	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(cost),
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
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)

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
