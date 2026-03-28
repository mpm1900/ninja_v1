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
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				result := game.GetAccuracyResult(g, *context.SourceActorID, config.Accuracy)
				if result.Success {
					damages := mutations.NewDamage(config, game.NewDamageConfig())
					transactions = append(
						transactions,
						mutations.MakeDamageTransactions(context, damages)...,
					)
				} else {
					transaction := game.MakeTransaction(mutations.AddLogs("Attack missed."), context)
					transactions = append(transactions, transaction)
				}

				return transactions
			},
		},
	}
}
