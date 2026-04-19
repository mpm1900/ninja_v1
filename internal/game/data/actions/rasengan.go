package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Rasengan = MakeRasengan()

func MakeRasengan() game.Action {
	ID := uuid.MustParse("054eb97a-cd6f-4428-8f54-96d9b6b33bfa")
	config := game.ActionConfig{
		Name:        "Rasengan",
		Nature:      game.Ptr(game.NsPure),
		Accuracy:    game.Ptr(90),
		Power:       game.Ptr(90),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(50),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return makeBasicAttack(ID, config)
}
