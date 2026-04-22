package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var MajesticFlame = MakeMajesticFlame()

func MakeMajesticFlame() game.Action {
	ID := uuid.MustParse("19c9f58e-9012-417f-af58-1c09d448f0dc")

	config := game.ActionConfig{
		Name:        "Majestic Flame",
		Description: "40% chance to burn targets.",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(100),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(120),
		Cooldown:    game.Ptr(2),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	action := makeBasicAttackWith(
		ID,
		config,
		func(g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			targets := g.GetTargets(context)
			for _, target := range targets {
				transactions = append(transactions, chanceBurn(config, target, 20)...)
			}

			return transactions
		},
		nil,
	)

	return action
}
