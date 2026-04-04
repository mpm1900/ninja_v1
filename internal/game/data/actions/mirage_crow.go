package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var MirageCrow = MakeMirageCrow()

func MakeMirageCrow() game.Action {
	ID := uuid.New()
	nature := game.NsYin
	targetCount := 1
	chakraCost := 30

	config := game.ActionConfig{
		Name:        "Mirage Crow",
		Description: "Lowers the target's Chakra Attack stat by 2 stages. User then switches out.",
		Nature:      &nature,
		TargetCount: &targetCount,
		Cost:        &chakraCost,
		Jutsu:       game.Genjutsu,
	}

	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				targets := g.GetTargets(context)
				for _, target := range targets {
					mut_ctx := game.Context{
						SourcePlayerID: &target.PlayerID,
						SourceActorID:  &target.ID,
						ParentActorID:  &target.ID,
					}
					mutation := mutations.AddModifiers(modifiers.ChakraAttackDown2Source)
					transaction := game.MakeTransaction(mutation, mut_ctx)
					transactions = append(transactions, transaction)
				}

				switch_mux := game.RemovePositions
				switch_ctx := game.NewContext()
				switch_ctx.TargetActorIDs = append(switch_ctx.TargetActorIDs, *context.SourceActorID)
				switch_tx := game.MakeTransaction(switch_mux, switch_ctx)
				transactions = append(transactions, switch_tx)

				return transactions
			},
		},
	}
}
