package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Chidori = MakeChidori()

func MakeChidori() game.Action {
	ID := uuid.MustParse("c1502330-764c-56f8-9c9e-f41b933a90f0")
	config := game.ActionConfig{
		Name:        "Chidori",
		Nature:      game.Ptr(game.NsLightning),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(120),
		Stat:        game.Ptr(game.StatChakraAttack),
		Cost:        game.Ptr(50),
		TargetCount: game.Ptr(1),
		Recoil:      game.Ptr(0.3),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}
	return makeBasicAttack(ID, config)
}
