package actions

import (
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"

	"github.com/google/uuid"
)

var Fireball = MakeFireball()

func MakeFireball() game.Action {
	accuracy := 80
	power := 50
	nature := game.NsFire
	stat := game.AttackNinjutsu
	config := game.ActionConfig{
		Name:     "火遁: Fireball",
		Nature:   &nature,
		Accuracy: &accuracy,
		Power:    &power,
		Stat:     &stat,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter, game.AliveFilter),
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: 0,
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
