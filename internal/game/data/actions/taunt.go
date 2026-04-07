package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Taunt = MakeTaunt()

func MakeTaunt() game.Action {
	nature := game.NsPure
	config := game.ActionConfig{
		Name:        "Taunt",
		Nature:      &nature,
		Jutsu:       game.Ninjutsu,
		Description: "Forces target to use only attacking moves.",
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.ComposeAF(game.ActiveFilter, game.TargetableFilter),
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				targets := g.GetTargets(context)
				if len(targets) == 1 {
					context.ParentActorID = &targets[0].ID
				}
				mutation := mutations.AddModifiers(true, modifiers.Taunted)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
