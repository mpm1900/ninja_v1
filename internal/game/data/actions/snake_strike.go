package actions

import (
	"math/rand/v2"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var SnakeStrike = MakeSnakeStrike()

func MakeSnakeStrike() game.Action {
	ID := uuid.MustParse("62587c38-1644-4910-a2c0-c44a6b27c576")

	config := game.ActionConfig{
		Name:        "Snake Strike",
		Description: "30% chance to paralyze, poison, or put target to sleep.",
		Nature:      game.Ptr(game.NsYang),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(60),
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
				roll := rand.IntN(3)
				switch roll {
				case 0:
					transactions = append(transactions, chanceParalysis(config, context, target, 30)...)
				case 1:
					transactions = append(transactions, chancePoison(config, context, target, 30)...)
				case 2:
					transactions = append(transactions, chanceSleep(config, context, target, 30)...)
				}
			}

			return transactions
		},
		nil,
	)
}
