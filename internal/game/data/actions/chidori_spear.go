package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var ChidoriSpear = MakeChidoriSpear()

func MakeChidoriSpear() game.Action {
	ID := uuid.MustParse("f89e4aba-35c8-4ddf-ab0b-bb809d5deb69")

	config := game.ActionConfig{
		Name:        "Chidori Spear",
		Description: "20% chance to paralyze the target.",
		Nature:      game.Ptr(game.NsLightning),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(90),
		Stat:        game.Ptr(game.StatChakraAttack),
		Cost:        game.Ptr(50),
		TargetCount: game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(getCriticalStage(0)),
		CritMod:     1.5,
	}

	return makeBasicAttackWith(
		ID,
		config,
		func(g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			targets := g.GetTargets(context)
			for _, target := range targets {
				transactions = append(transactions, chanceParalysis(config, target, 20)...)
			}

			return transactions
		},
		nil,
	)
}
