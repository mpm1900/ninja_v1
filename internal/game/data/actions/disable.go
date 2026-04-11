package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Disable = MakeDisable()

func MakeDisable() game.Action {
	cooldown := 1
	nature := game.NsYin
	config := game.ActionConfig{
		Name:        "Disable",
		Nature:      &nature,
		Cooldown:    &cooldown,
		Jutsu:       game.Genjutsu,
		Description: "Disable's the target's last used action.",
	}
	return game.Action{
		ID:              uuid.MustParse("5cf69985-6785-56a6-b879-e02cb6207960"),
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
						mutation := mutations.AddModifiers(true, modifiers.Disabled(*target.LastUsedActionID))
						transaction := game.MakeTransaction(mutation, mut_ctx)
						transactions = append(transactions, transaction)
					}
				}

				return transactions
			},
		},
	}
}
