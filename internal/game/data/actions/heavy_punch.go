package actions

import (
	"math/rand"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var HeavyPunch = MakeHeavyPunch()

func MakeHeavyPunch() game.Action {
	ID := uuid.MustParse("420bad58-1238-4124-909e-09ef76d743e8")
	config := game.ActionConfig{
		Name:        "Heavy Punch",
		Description: "30% chance to paralyze the target.",
		Nature:      game.Ptr(game.NsTai),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(80),
		Stat:        game.Ptr(game.StatAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(0),
		Cooldown:    game.Ptr(0),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return makeBasicAttackWith(ID, config, func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction {
		targets := g.GetTargets(context)
		for _, target := range targets {
			// on 30% chance
			roll := rand.Intn(100)
			if roll > 30 {
				continue
			}
			transactions = append(transactions, applyParalysis(config, target)...)
		}

		return transactions
	}, nil)
}
