package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var LightningHound = MakeLightningHound()

func MakeLightningHound() game.Action {
	ID := uuid.MustParse("5a35a0b0-160a-4b73-9b81-bb301e7c8f7e")

	config := game.ActionConfig{
		Name:        "Lightning Hound",
		Description: "10% chance to paralyze target.",
		Nature:      game.Ptr(game.NsLightning),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(95),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(50),
		Cooldown:    game.Ptr(0),
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
				transactions = append(transactions, chanceParalysis(config, context, target, 10)...)
			}

			return transactions
		},
		nil,
	)
}
