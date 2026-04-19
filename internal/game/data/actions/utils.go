package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

func makeBasicAttackWith(
	ID uuid.UUID,
	config game.ActionConfig,
	with func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction,
	before func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction,
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

				if before != nil {
					transactions = before(g, context, transactions)
				}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damages := game.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor()))
				transactions = append(
					transactions,
					game.MakeDamageTransactions(context, damages)...,
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
	return makeBasicAttackWith(ID, config, nil, nil)
}

func applyBurn(actor game.Actor) []game.GameTransaction {
	transactions := []game.GameTransaction{}

	mut_ctx := game.Context{
		SourcePlayerID: &actor.PlayerID,
		SourceActorID:  &actor.ID,
		ParentActorID:  nil, // do not remove on switch
		TargetActorIDs: []uuid.UUID{actor.ID},
	}

	mod := mutations.AddStatus(true, modifiers.Burned)
	mod_tx := game.MakeTransaction(mod, mut_ctx)

	mut := mutations.Burn
	mut_tx := game.MakeTransaction(mut, mut_ctx)
	transactions = append(transactions, mod_tx, mut_tx)

	return transactions
}

func applyParalysis(actor game.Actor) []game.GameTransaction {
	transactions := []game.GameTransaction{}

	mut_ctx := game.Context{
		SourcePlayerID: &actor.PlayerID,
		SourceActorID:  &actor.ID,
		ParentActorID:  nil, // do not remove on switch
		TargetActorIDs: []uuid.UUID{actor.ID},
	}

	mod := mutations.AddStatus(true, modifiers.Paralysis)
	mod_tx := game.MakeTransaction(mod, mut_ctx)

	mut := mutations.Paralyze
	mut_tx := game.MakeTransaction(mut, mut_ctx)
	transactions = append(transactions, mod_tx, mut_tx)

	return transactions
}
