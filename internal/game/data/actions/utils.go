package actions

import (
	"fmt"
	"math/rand/v2"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

func makeBasicAttackWith(
	ID uuid.UUID,
	config game.ActionConfig,
	onSuccess func(g game.Game, context game.Context) []game.GameTransaction,
	before func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction,
) game.Action {
	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            modifiers.UseStaminaCost(*config.Cost),
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

				action_config, _ := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(action_config)
				dmg_config := game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor())
				if onSuccess != nil {
					dmg_config.OnSuccess = onSuccess
				}
				damages := game.NewDamage(action_config, dmg_config)
				transactions = append(
					transactions,
					game.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}

func makeBasicAttack(ID uuid.UUID, config game.ActionConfig) game.Action {
	return makeBasicAttackWith(ID, config, nil, nil)
}

func applyStatus(config game.ActionConfig, actor game.Actor, modifier game.Modifier, mutation game.GameMutation) []game.GameTransaction {
	transactions := []game.GameTransaction{}

	if mutations.CheckJutsuImmunity(config, actor) {
		log_ctx := game.MakeContextForActor(actor)
		log := game.NewLogContext(fmt.Sprintf("| $source$ was immune to %s.", config.Jutsu), log_ctx)
		tx := game.AddLogs(log)
		transactions = append(transactions, game.MakeTransaction(tx, log_ctx))

		return transactions
	}

	ctx := game.MakeContextForActor(actor)
	ctx.ParentActorID = nil // do not remove on switch

	mod := mutations.AddStatus(true, modifier)
	mod_tx := game.MakeTransaction(mod, ctx)

	mut_tx := game.MakeTransaction(mutation, ctx)
	transactions = append(transactions, mod_tx, mut_tx)

	return transactions
}

func applyBurn(config game.ActionConfig, actor game.Actor) []game.GameTransaction {
	return applyStatus(config, actor, modifiers.Burned, mutations.Burn)
}
func chanceBurn(config game.ActionConfig, actor game.Actor, chance int) []game.GameTransaction {
	roll := rand.IntN(100)
	if roll > chance {
		return []game.GameTransaction{}
	}

	return applyBurn(config, actor)
}

func applyParalysis(config game.ActionConfig, actor game.Actor) []game.GameTransaction {
	return applyStatus(config, actor, modifiers.Paralysis, mutations.Paralyze)
}
func chanceParalysis(config game.ActionConfig, actor game.Actor, chance int) []game.GameTransaction {
	roll := rand.IntN(100)
	if roll > chance {
		return []game.GameTransaction{}
	}

	return applyParalysis(config, actor)
}

func applySleep(config game.ActionConfig, actor game.Actor) []game.GameTransaction {
	return applyStatus(config, actor, modifiers.Sleeping, mutations.Sleep)
}

func applySummon(context game.Context, def game.ActorDef, actions []game.Action) []game.GameTransaction {
	transactions := []game.GameTransaction{}

	mut := game.GameMutation{
		Delta: func(mp, mg game.Game, mc game.Context) game.Game {
			mg.UpdateActor(*mc.SourceActorID, func(a game.Actor) game.Actor {
				summon := game.MakeActor(
					def,
					a.PlayerID,
					a.Experience,
					nil,
					nil,
					append(actions, game.CancelSummon),
					game.FocusNone,
					map[game.ActorStat]int{},
				)
				a.SetSummonFromActor(&summon, false)
				return a
			})
			mg.UpdatePlayer(*mc.SourcePlayerID, func(p game.Player) game.Player {
				p.UsedSummon = true
				return p
			})
			return mg
		},
	}

	transactions = append(
		transactions,
		game.MakeTransaction(mut, context),
	)

	return transactions
}

func checkPlayerHasModifier(g game.Game, context game.Context, modifierID uuid.UUID) bool {
	for _, tx := range g.GetModifiers() {
		if tx.Context.SourcePlayerID == nil {
			continue
		}

		if *tx.Context.SourcePlayerID == *context.SourcePlayerID && tx.Mutation.ID == modifierID {
			return true
		}

	}

	return false
}
