package modifiers

import (
	"fmt"
	"math/rand/v2"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"
)

func ApplyModifier(config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier) []game.GameTransaction {
	transactions := []game.GameTransaction{}

	if mutations.CheckJutsuImmunity(config, actor) {
		log_ctx := game.MakeContextForActor(actor)
		log := game.NewLogContext(fmt.Sprintf("| $source$ was immune to %s.", config.Jutsu), log_ctx)
		tx := game.AddLogs(log)
		transactions = append(transactions, game.MakeTransaction(tx, log_ctx))

		return transactions
	}

	ctx := game.MakeContextForActor(actor)
	ctx.ModifierID = modifier.GroupID
	mod := mutations.AddModifiers(true, modifier)
	mod_tx := game.MakeTransaction(mod, ctx)
	transactions = append(transactions, mod_tx)

	return transactions
}
func ChanceModifier(config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier, chance int) []game.GameTransaction {
	roll := rand.IntN(100)
	if roll > chance {
		return []game.GameTransaction{}
	}

	return ApplyModifier(config, context, actor, modifier)
}

func ApplyStatus(config game.ActionConfig, context game.Context, actor game.Actor, modifier game.Modifier, mutation game.GameMutation) []game.GameTransaction {
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

func ApplyBurn(config game.ActionConfig, context game.Context, actor game.Actor) []game.GameTransaction {
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
