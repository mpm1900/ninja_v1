package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var DragonFire = MakeDragonFire()

func MakeDragonFire() game.Action {
	ID := uuid.MustParse("dca159df-75fb-4cf0-85c5-db69d987a029")

	config := game.ActionConfig{
		Name:        "Dragon Fire",
		Description: "5% chance to burn target. Never Misses.",
		Nature:      game.Ptr(game.NsFire),
		Power:       game.Ptr(80),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(0),
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
				transactions = append(transactions, chanceBurn(config, target, 5)...)
			}

			return transactions
		},
		nil,
	)
}
