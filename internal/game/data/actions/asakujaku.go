package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Asakujaku = MakeAsakujaku()

func MakeAsakujaku() game.Action {
	ID := uuid.MustParse("5c7a660c-7d53-48a2-94c3-c2b683fed948")

	config := game.ActionConfig{
		Name:        "6th Gate: Asakujaku",
		Description: "Deals 30% recoil damage. 10% chance to burn target. High crit chance.",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(95),
		Power:       game.Ptr(120),
		Stat:        game.Ptr(game.StatAttack),
		TargetCount: game.Ptr(1),
		Recoil:      game.Ptr(0.3),
		Cost:        game.Ptr(0),
		Cooldown:    game.Ptr(0),
		Jutsu:       game.Taijutsu,
		CritChance:  game.Ptr(15),
		CritMod:     1.5,
	}

	return makeBasicAttackWith(
		ID,
		config,
		func(g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			targets := g.GetTargets(context)
			for _, target := range targets {
				transactions = append(transactions, chanceBurn(config, target, 10)...)
			}

			return transactions
		},
		nil,
	)
}
