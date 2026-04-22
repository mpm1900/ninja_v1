package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var FalseDarkness = MakeFalseDarkness()

func MakeFalseDarkness() game.Action {
	ID := uuid.MustParse("99338b50-de10-4747-9e41-847677db4ca0")

	config := game.ActionConfig{
		Name:        "False Darkness",
		Description: "Grants the user Lightning nature until end of turn.",
		Nature:      game.Ptr(game.NsLightning),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(95),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(getCriticalStage(0)),
		CritMod:     1.5,
	}

	action := makeBasicAttackWith(ID, config, nil, func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction {
		add_mut := mutations.AddModifiers(false, modifiers.AddNature(game.NsLightning, 0))
		add_tx := game.MakeTransaction(add_mut, context)
		transactions = append(transactions, add_tx)
		return transactions
	})
	return action
}
