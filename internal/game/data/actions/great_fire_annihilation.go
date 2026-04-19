package actions

import (
	"fmt"
	"math/rand"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var GreatFireAnnihilation = MakeGreatFireAnnihilation()

func MakeGreatFireAnnihilation() game.Action {
	ID := uuid.MustParse("d97ee3bb-7afa-47de-9f8d-2ee77ba6dfe6")

	config := game.ActionConfig{
		Name:        "Great Fire Annihilation",
		Description: "30% chance to burn target.",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(100),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(1),
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
			fmt.Println("BURN! roll=", roll)
			transactions = append(transactions, applyBurn(target)...)
		}

		return transactions
	}, nil)
}
