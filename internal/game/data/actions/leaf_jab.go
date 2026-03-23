package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

func MakeLeafJab() game.Action {
	config := game.ActionConfig{
		Name: "Leaf Jab",
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetPredicate: game.AllFilter,
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.MakeActionMutation(
			0,
			game.AllGameFilter,
			func(g game.Game, context *game.Context) []game.Transaction[game.GameMutation] {
				return []game.Transaction[game.GameMutation]{
					game.MakeTransaction(mutations.NewDamage(game.StatTaijutsu, 50), context),
				}
			},
		),
	}
}
