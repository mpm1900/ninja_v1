package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Coercion = MakeCoercion()

func MakeCoercion() game.Action {
	cooldown := 1
	nature := game.NsYin
	config := game.ActionConfig{
		Name:     "Sharingan: Coercion",
		Nature:   &nature,
		Cooldown: &cooldown,
		Jutsu:    game.Genjutsu,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherTeamFilter, game.ActiveFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityProtect,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
				game.SourceIsActionOnCooldown,
				game.SourceHasActiveTurns(1),
			),
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				mutation := mutations.AddModifiers(modifiers.Stunned)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
