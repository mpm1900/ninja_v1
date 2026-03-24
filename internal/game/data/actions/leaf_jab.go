package actions

import (
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"

	"github.com/google/uuid"
)

func MakeLeafJab() game.Action {
	accuracy := 100
	config := game.ActionConfig{
		Name:     "Leaf Jab",
		Accuracy: &accuracy,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.AliveFilter),
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: 0,
			Filter:   game.AllGameFilter,
			Delta: func(g game.Game, context *game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				// accuracy checks TODO
				// fmt.Print(config.Accuracy)

				transactions = append(
					transactions,
					game.MakeTransaction(
						mutations.NewDamage(game.AttackTaijutsu, 50, config.Nature),
						context,
					),
				)

				return transactions
			},
		},
	}
}
