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
		Name:        "Coercion",
		Nature:      &nature,
		Cooldown:    &cooldown,
		Jutsu:       game.Genjutsu,
		Description: "Forces the target to use only their last used action.",
	}
	return game.Action{
		ID:              uuid.MustParse("06840403-52cc-4e8a-95eb-318cf012e634"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPrioritySlow,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
				game.SourceIsActionOffCooldown,
				game.SourceHasActiveTurns(1),
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				targets := g.GetTargets(context)
				for _, target := range targets {
					mut_ctx := game.MakeContextForActor(target)
					if target.LastUsedActionID != nil {
						mutation := mutations.AddModifiers(true, modifiers.Coerced(*target.LastUsedActionID))
						transaction := game.MakeTransaction(mutation, mut_ctx)
						transactions = append(transactions, transaction)
					}
				}

				return transactions
			},
		},
	}
}
