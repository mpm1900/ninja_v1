package actions

import (
	"fmt"
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"

	"github.com/google/uuid"
)

type GameTransaction = game.Transaction[game.GameMutation]
type ActionMutation = game.Mutation[game.Game, []GameTransaction]

func MakeLeafJab() game.Action {
	accuracy := 100
	config := game.ActionConfig{
		Name:     "Leaf Jab",
		Accuracy: &accuracy,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetPredicate: game.OtherFilter,
		ContextValidate: game.TargetLengthFilter(1),
		Mutation: ActionMutation{
			Priority: 0,
			Filter:   game.AllGameFilter,
			Delta: func(g game.Game, context *game.Context) []GameTransaction {
				transactions := []GameTransaction{}

				// accuracy checks
				fmt.Print(config.Accuracy)

				transactions = append(
					transactions,
					game.MakeTransaction(
						mutations.NewDamage(game.AttackTaijutsu, 50),
						context,
					),
				)

				return transactions
			},
		},
	}
}
