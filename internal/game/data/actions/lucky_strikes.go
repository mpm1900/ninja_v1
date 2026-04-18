package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var LuckyStrikes = MakeLuckyStrikes()

func MakeLuckyStrikes() game.Action {
	accuracy := 80
	power := 10
	stat := game.StatAttack
	nature := game.NsTai
	chakraCost := 30
	config := game.ActionConfig{
		Name:     "Lucky Strikes",
		Accuracy: &accuracy,
		Power:    &power,
		Stat:     &stat,
		Nature:   &nature,
		Cost:     &chakraCost,
		Jutsu:    game.Taijutsu,
	}
	return game.Action{
		ID:              uuid.MustParse("4ac4894c-2ff3-5142-b087-a8924837cefc"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		Cost:            mutations.UseStaminaSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				damage_config := game.NewDamageConfig(1, 1)
				damage_config.Repeat = true
				damage_config.RepeatMax = -1
				damages := mutations.NewDamage(conf, damage_config)
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
