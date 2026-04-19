package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

func makeBasicAttackWith(
	ID uuid.UUID,
	config game.ActionConfig,
	with func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction,
) game.Action {
	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(*config.Cost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
				game.SourceIsActionOffCooldown,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor()))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				if with == nil {
					return transactions
				}

				return with(g, context, transactions)
			},
		},
	}
}

func makeBasicAttack(ID uuid.UUID, config game.ActionConfig) game.Action {
	return makeBasicAttackWith(ID, config, nil)
}
