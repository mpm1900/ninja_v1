package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var BodyReplacement = MakeBodyReplacement()

func MakeBodyReplacement() game.Action {
	config := game.ActionConfig{
		Name:        "Body Replacement",
		Nature:      game.Ptr(game.NsYin),
		Cooldown:    game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		Description: "Protects the user from actions this turn. 1 turn cooldown.",
	}
	return game.Action{
		ID:              uuid.MustParse("d3765608-4b30-5c4c-b5a9-f4132f0bbb7c"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityProtect,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
				game.SourceIsActionOffCooldown,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				mutation := mutations.AddModifiers(false, modifiers.Protected)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
