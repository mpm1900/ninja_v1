package actions

import (
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"

	"github.com/google/uuid"
)

var Recover = MakeRecover()

func MakeRecover() game.Action {
	nature := game.NsYang
	stat := game.AttackNinjutsu
	targetCount := 1
	chakraCost := 30
	config := game.ActionConfig{
		Name:        "Recover",
		Nature:      &nature,
		Stat:        &stat,
		TargetCount: &targetCount,
		Cost:        &chakraCost,
	}

	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.ComposeAF(game.ActiveFilter),
		ContextValidate: game.TargetLengthFilter(*config.TargetCount),
		Cost:            mutations.UseChakraSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				heal := mutations.NewHeal(config, 0.5)
				transactions = append(
					transactions,
					game.MakeTransaction(heal, context),
				)

				return transactions
			},
		},
	}
}
