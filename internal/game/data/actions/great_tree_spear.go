package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var GreatTreeSpear = MakeGreatTreeSpear()

func MakeGreatTreeSpear() game.Action {
	ID := uuid.New()
	accuracy := 70
	power := 120
	nature := game.NsWood
	stat := game.ChakraAttack
	targetCount := 1
	chakraCost := 90

	config := game.ActionConfig{
		Name:        "Great Tree Spear",
		Nature:      &nature,
		Accuracy:    &accuracy,
		Power:       &power,
		Stat:        &stat,
		TargetCount: &targetCount,
		Cost:        &chakraCost,
		Jutsu:       game.Ninjutsu,
	}

	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
				game.SourceIsActionOffCooldown,
			),
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				damages := mutations.NewDamage(config, game.NewDamageConfig(1, 1))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
