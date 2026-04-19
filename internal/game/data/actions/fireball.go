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
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(70),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return makeBasicAttack(ID, config)
}
