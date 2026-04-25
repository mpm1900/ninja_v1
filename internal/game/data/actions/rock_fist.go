package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var RockFist = MakeRockFist()

func MakeRockFist() game.Action {
	ID := uuid.MustParse("6e015595-18df-4f2b-b4da-cf97863a3f4e")
	config := game.ActionConfig{
		Name:        "Rock Fist",
		Description: "30% chance to paralyze the target.",
		Nature:      game.Ptr(game.NsEarth),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(80),
		Stat:        game.Ptr(game.StatAttack),
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
		func(g game.Game, _, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			targets := g.GetTargets(context)
			for _, target := range targets {
				transactions = append(transactions, chanceParalysis(config, context, target, 30)...)
			}

			return transactions
		},
		nil,
	)
}
