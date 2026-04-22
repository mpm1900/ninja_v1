package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var PatternBreak = MakePatternBreak()

func MakePatternBreak() game.Action {
	config := game.ActionConfig{
		Name:        "Pattern Break",
		Nature:      game.Ptr(game.NsYin),
		Jutsu:       game.Ninjutsu,
		Description: "Target cannot repeat actions.",
	}
	return game.Action{
		ID:              uuid.MustParse("c62f29ad-2f3e-5e5e-b045-bb0ed58837bc"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				targets := g.GetTargets(context)
				for _, target := range targets {
					context.ParentActorID = &target.ID
					mod := modifiers.PatternBroke
					mutation := mutations.AddModifiers(false, mod)
					transaction := game.MakeTransaction(mutation, context)
					transactions = append(transactions, transaction)
				}

				return transactions
			},
		},
	}
}
