package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Amaterasu = MakeAmaterasu()

func MakeAmaterasu() game.Action {
	ID := uuid.MustParse("d103e605-9381-52fd-9cb8-450b7315a9f9")

	config := game.ActionConfig{
		Name:        "Amaterasu",
		Description: "Burns target.",
		Nature:      game.Ptr(game.NsYin),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Power:       game.Ptr(20),
		Accuracy:    game.Ptr(100),
		Cost:        game.Ptr(30),
		Jutsu:       game.Genjutsu,
		CritChance:  game.Ptr(getCriticalStage(0)),
		CritMod:     1.5,
	}

	action := makeBasicAttackWith(
		ID,
		config,
		func(g game.Game, _, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			targets := g.GetTargets(context)
			for _, target := range targets {
				transactions = append(transactions, modifiers.ApplyBurn(config, context, target)...)
			}

			return transactions
		},
		nil,
	)

	return action
}
