package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var BodyPress = MakeBodyPress()

func MakeBodyPress() game.Action {
	accuracy := 100
	power := 80
	stat := game.StatDefense
	nature := game.NsTai
	chakraCost := 0
	config := game.ActionConfig{
		Name:        "Body Press",
		Description: "Damage is based of the user's Defense stat rather than Attack.",
		Accuracy:    &accuracy,
		Power:       &power,
		Stat:        &stat,
		Nature:      &nature,
		Cost:        &chakraCost,
		Jutsu:       game.Taijutsu,
	}
	return game.Action{
		ID:              uuid.MustParse("05b5376a-5c76-4f72-bc2c-c148ad068e40"),
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
