package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var BodyFlicker = MakeBodyFlicker()

func MakeBodyFlicker() game.Action {
	ID := uuid.MustParse("f052f07c-bb06-4f44-8b26-ec2f17401446")
	nature := game.NsTai
	targetCount := 1

	config := game.ActionConfig{
		Name:        "Body Flicker",
		Description: "User then switches out.",
		Nature:      &nature,
		TargetCount: &targetCount,
		Jutsu:       game.Taijutsu,
		Power:       game.Ptr(70),
		Accuracy:    game.Ptr(100),
		Stat:        game.Ptr(game.StatAttack),
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, 1))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

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
