package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Glare = MakeGlare()

func MakeGlare() game.Action {
	cooldown := 1
	nature := game.NsYin
	config := game.ActionConfig{
		Name:        "Sharingan: Glare",
		Nature:      &nature,
		Cooldown:    &cooldown,
		Jutsu:       game.Genjutsu,
		Description: "Stuns the target this turn. Fails unless it is the user's first turn switched in.",
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityP3,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
				game.SourceIsActionOffCooldown,
				game.SourceHasActiveTurns(1),
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				mutation := mutations.AddModifiers(true, modifiers.Stunned)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
