package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var LeafJab = MakeLeafJab()

func MakeLeafJab() game.Action {
	accuracy := 100
	power := 90
	stat := game.Attack
	nature := game.NsWood
	chakraCost := 30
	config := game.ActionConfig{
		Name:     "Leaf Jab",
		Accuracy: &accuracy,
		Power:    &power,
		Stat:     &stat,
		Nature:   &nature,
		Cost:     &chakraCost,
		Jutsu:    game.Taijutsu,
	}
	return game.Action{
		ID:              uuid.MustParse("b23ace96-eb09-5bf7-b884-7ef8e8fc544d"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		Cost:            mutations.UseStaminaSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityP1,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
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
