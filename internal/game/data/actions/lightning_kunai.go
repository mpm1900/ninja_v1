package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var LightningKunai = MakeLightningKunai()

func MakeLightningKunai() game.Action {
	ID := uuid.MustParse("5b5a4fe7-5c3c-4279-ae28-bfde927f8d8b")

	config := game.ActionConfig{
		Name:        "Lightning Kunai",
		Description: "10% chance to paralyze target.",
		Nature:      game.Ptr(game.NsLightning),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(55),
		Stat:        game.Ptr(game.StatAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(40),
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
				transactions = append(transactions, modifiers.ChanceParalysis(config, context, target, 10)...)
			}

			return transactions
		},
		nil,
	)
}
