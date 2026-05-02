package modifiers

import (
	"fmt"
	"math/rand/v2"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

func RemoveModifierSource(id uuid.UUID, actor game.Actor) []game.GameTransaction {
	transactions := []game.GameTransaction{}

	ctx := game.MakeContextForActor(actor)
	ctx.ModifierID = &id
	mod := mutations.RemoveModifierWhere(func(t game.Transaction[game.Modifier]) bool {
		if t.Context.SourceActorID == nil {
			return false
		}

		return actor.ID == *t.Context.SourceActorID &&
			(t.Mutation.ID == id || (t.Mutation.GroupID != nil && *t.Mutation.GroupID == id))
	})
	mod_tx := game.MakeTransaction(mod, ctx)
	transactions = append(transactions, mod_tx)

	return transactions
}
func applyModifier(checkWarded bool, config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier) []game.GameTransaction {
	transactions := []game.GameTransaction{}

	if mutations.CheckJutsuImmunity(config, actor) {
		log_ctx := game.MakeContextForActor(actor)
		log := game.MakeGameLog(fmt.Sprintf("$source$ was immune to %s.", config.Jutsu), log_ctx, 1)
		tx := game.AddLogs(log)
		transactions = append(transactions, game.MakeTransaction(tx, log_ctx))

		return transactions
	}

	ctx := game.MakeContextForActor(actor)
	ctx.ModifierID = modifier.GroupID
	mod := mutations.AddModifiers(checkWarded, modifier)
	mod_tx := game.MakeTransaction(mod, ctx)
	transactions = append(transactions, mod_tx)

	return transactions
}
func ApplyModifier(config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier) []game.GameTransaction {
	return applyModifier(true, config, context, actor, modifier)
}
func ApplyModifierBypass(config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier) []game.GameTransaction {
	return applyModifier(false, config, context, actor, modifier)
}
func ChanceModifier(config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier, chance int) []game.GameTransaction {
	roll := rand.IntN(100)
	if roll > chance {
		return []game.GameTransaction{}
	}

	return ApplyModifier(config, context, actor, modifier)
}

func applyStatus(checkWarded bool, config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier, mutation game.GameMutation) []game.GameTransaction {
	transactions := []game.GameTransaction{}

	if mutations.CheckJutsuImmunity(config, actor) {
		log_ctx := game.MakeContextForActor(actor)
		log := game.MakeGameLog(fmt.Sprintf("$source$ was immune to %s.", config.Jutsu), log_ctx, 1)
		tx := game.AddLogs(log)
		transactions = append(transactions, game.MakeTransaction(tx, log_ctx))

		return transactions
	}
	if mutations.CheckImmunity(modifier.ID, actor) {
		log_ctx := game.MakeContextForActor(actor)
		log := game.MakeGameLog("$source$ was immune.", log_ctx, 1)
		tx := game.AddLogs(log)
		transactions = append(transactions, game.MakeTransaction(tx, log_ctx))

		return transactions
	}

	ctx := game.MakeContextForActor(actor)
	ctx.ParentActorID = nil // do not remove on switch

	mod := mutations.AddStatus(checkWarded, modifier)
	mod_tx := game.MakeTransaction(mod, ctx)

	mut_tx := game.MakeTransaction(mutation, ctx)
	transactions = append(transactions, mod_tx, mut_tx)

	return transactions
}
func ApplyStatus(config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier, mutation game.GameMutation) []game.GameTransaction {
	return applyStatus(true, config, context, actor, modifier, mutation)
}
func ApplyStatusBypass(config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier, mutation game.GameMutation) []game.GameTransaction {
	return applyStatus(false, config, context, actor, modifier, mutation)
}

func ApplyBurn(config game.ActionConfig, context game.Context, actor game.Actor) []game.GameTransaction {
	_, ok := actor.Natures[game.NsFire]
	if ok {
		// fire nature immune to burn
		return []game.GameTransaction{}
	}

	return ApplyStatus(config, context, actor, Burned, mutations.Burn)
}
func ChanceBurn(config game.ActionConfig, context game.Context, actor game.Actor, chance int) []game.GameTransaction {
	roll := rand.IntN(100)
	if roll > chance {
		return []game.GameTransaction{}
	}

	return ApplyBurn(config, context, actor)
}

func ApplyParalysis(config game.ActionConfig, context game.Context, actor game.Actor) []game.GameTransaction {
	_, ok := actor.Natures[game.NsLightning]
	if ok {
		// lightning nature immune to paralysis
		return []game.GameTransaction{}
	}
	return ApplyStatus(config, context, actor, Paralysis, mutations.Paralyze)
}
func ChanceParalysis(config game.ActionConfig, context game.Context, actor game.Actor, chance int) []game.GameTransaction {
	roll := rand.IntN(100)
	if roll > chance {
		return []game.GameTransaction{}
	}

	return ApplyParalysis(config, context, actor)
}

func ApplySleep(config game.ActionConfig, context game.Context, actor game.Actor) []game.GameTransaction {
	return ApplyStatus(config, context, actor, Sleeping, mutations.Sleep)
}
func ChanceSleep(config game.ActionConfig, context game.Context, actor game.Actor, chance int) []game.GameTransaction {
	roll := rand.IntN(100)
	if roll > chance {
		return []game.GameTransaction{}
	}

	return ApplySleep(config, context, actor)
}

func ApplyPoison(config game.ActionConfig, context game.Context, actor game.Actor) []game.GameTransaction {
	return ApplyStatus(config, context, actor, Poisoned, mutations.Poison)
}
func ChancePoison(config game.ActionConfig, context game.Context, actor game.Actor, chance int) []game.GameTransaction {
	roll := rand.IntN(100)
	if roll > chance {
		return []game.GameTransaction{}
	}

	return ApplyPoison(config, context, actor)
}

func ApplyImmunity(modifier_id uuid.UUID, immunity_id uuid.UUID) game.ActorMutation {
	return game.MakeActorMutation(
		&modifier_id,
		game.MutPriorityDefault,
		game.SourceFilter,
		func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
			if a.Immunities == nil {
				a.Immunities = map[uuid.UUID]struct{}{}
			}
			a.Immunities[immunity_id] = struct{}{}
			return a
		})
}
