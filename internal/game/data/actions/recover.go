package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Recover = MakeRecover()

func MakeRecover() game.Action {
	nature := game.NsYang
	targetCount := 1
	chakraCost := 30
	config := game.ActionConfig{
		Name:        "Recover",
		Nature:      &nature,
		TargetCount: &targetCount,
		Cost:        &chakraCost,
		Jutsu:       game.Senjutsu,
	}

	return game.Action{
		ID:              uuid.MustParse("c0756ddc-2611-5eef-82cc-c2bc03f9f01c"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.ComposeAF(game.ActiveFilter),
		ContextValidate: game.TargetLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf, _ := game.GetActiveActionConfig(g, config)
				heal := game.NewHeal(conf, 0.5)
				transactions = append(
					transactions,
					game.MakeTransaction(heal, context),
				)

				return transactions
			},
		},
	}
}
