package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Amaterasu = MakeAmaterasu()

func MakeAmaterasu() game.Action {
	ID := uuid.New()
	nature := game.NsYin
	targetCount := 1
	chakraCost := 30

	config := game.ActionConfig{
		Name:        "Amaterasu",
		Description: "Burns taraget.",
		Nature:      &nature,
		TargetCount: &targetCount,
		Cost:        &chakraCost,
		Jutsu:       game.Genjutsu,
	}

	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
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
						ParentActorID:  nil,
					}
					mutation := mutations.AddModifiers(true, modifiers.Burned)
					transaction := game.MakeTransaction(mutation, mut_ctx)
					transactions = append(transactions, transaction)
				}

				return transactions
			},
		},
	}
}
