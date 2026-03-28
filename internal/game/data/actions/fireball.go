package actions

import (
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"

	"github.com/google/uuid"
)

var Fireball = MakeFireball()

func MakeFireball() game.Action {
	accuracy := 50
	power := 50
	nature := game.NsFire
	stat := game.AttackNinjutsu
	targetCount := 1
	chakraCost := 30
	config := game.ActionConfig{
		Name:        "火遁: Fireball",
		Nature:      &nature,
		Accuracy:    &accuracy,
		Power:       &power,
		Stat:        &stat,
		TargetCount: &targetCount,
	}

	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseChakra(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				damages := mutations.NewDamage(config, game.NewDamageConfig())
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
