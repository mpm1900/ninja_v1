package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var GreatFireball = MakeGreatFireball()

func MakeGreatFireball() game.Action {
	ID := uuid.MustParse("57ddb3d7-0853-4a64-885b-18d93286c806")

	config := game.ActionConfig{
		Name:        "Great Fireball",
		Description: "20% chance to burn target.",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(90),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(60),
		Cooldown:    game.Ptr(0),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return makeBasicAttackWith(
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
}
