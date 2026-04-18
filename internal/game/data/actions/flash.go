package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Flash = MakeBlindingFlash()

func MakeBlindingFlash() game.Action {
	cooldown := 1
	nature := game.NsYin
	config := game.ActionConfig{
		Name:        "Flash",
		Nature:      &nature,
		Cooldown:    &cooldown,
		Jutsu:       game.Genjutsu,
		Description: "Stuns the target this turn. Fails unless it is the user's first turn switched in.",
	}
	return game.Action{
		ID:              uuid.MustParse("4cf69985-6785-56a6-b879-e02cb6207960"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
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

				mutation := mutations.AddModifiers(false, modifiers.Stunned)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
