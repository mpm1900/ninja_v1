package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Fireball = MakeFireball()

func MakeFireball() game.Action {
	ID := uuid.MustParse("aaf5174b-f386-54b1-84c4-0c062937c770")

	config := game.ActionConfig{
		Name:        "Fireball",
		Description: "10% chance to burn target.",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(70),
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
		func(g game.Game, _, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			targets := g.GetTargets(context)
			for _, target := range targets {
				transactions = append(transactions, chanceBurn(config, context, target, 10)...)
			}

			return transactions
		},
		nil,
	)
}
