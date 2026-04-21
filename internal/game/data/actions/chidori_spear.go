package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var ChidoriSpear = MakeChidoriSpear()

func MakeChidoriSpear() game.Action {
	ID := uuid.MustParse("d55c8221-fc03-4ae0-9737-cb5c7db88f73")

	config := game.ActionConfig{
		Name:        "Chidori Spear",
		Description: "20% chance to paralyze the target.",
		Nature:      game.Ptr(game.NsLightning),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(85),
		Stat:        game.Ptr(game.StatChakraAttack),
		Cost:        game.Ptr(50),
		TargetCount: game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return makeBasicAttackWith(ID, config, func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction {
		targets := g.GetTargets(context)
		for _, target := range targets {
			transactions = append(transactions, chanceParalysis(config, target, 20)...)
		}

		return transactions
	}, nil)
}
