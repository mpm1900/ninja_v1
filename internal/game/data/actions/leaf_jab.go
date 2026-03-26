package actions

import (
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"

	"github.com/google/uuid"
)

var LeafJab = MakeLeafJab()

func MakeLeafJab() game.Action {
	accuracy := 100
	power := 90
	stat := game.AttackTaijutsu
	nature := game.NsTai
	config := game.ActionConfig{
		Name:     "Leaf Jab",
		Accuracy: &accuracy,
		Power:    &power,
		Stat:     &stat,
		Nature:   &nature,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter, game.AliveFilter),
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityFast,
			Filter:   game.AllGameFilter,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				// accuracy checks TODO
				// fmt.Print(config.Accuracy)

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
