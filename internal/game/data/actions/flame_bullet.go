package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var FlameBullet = MakeFlameBullet()

func MakeFlameBullet() game.Action {
	ID := uuid.MustParse("aaf5174b-f386-54b1-84c4-0c062937c770")

	config := game.ActionConfig{
		Name:        "Flame Bullet",
		Description: "+1 priority.",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(60),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(0),
		Cooldown:    game.Ptr(0),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return makeBasicAttack(ID, config)
}
