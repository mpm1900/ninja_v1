package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var KusariChains = MakeKusariChains()

func MakeKusariChains() game.Action {
	ID := uuid.MustParse("65a8447d-4262-454d-a4ad-062993b1f8ad")

	config := game.ActionConfig{
		Name:        "Kusari Chains",
		Description: "30% chance to stun target.",
		Nature:      game.Ptr(game.NsTai),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(80),
		Stat:        game.Ptr(game.StatAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(0),
		Cooldown:    game.Ptr(0),
		Jutsu:       game.Bukijutsu,
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
				transactions = append(transactions, chanceModifier(config, target, modifiers.Stunned, 30)...)
			}

			return transactions
		},
		nil,
	)
}
