package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var LeafJab = MakeLeafJab()

func MakeLeafJab() game.Action {
	ID := uuid.MustParse("b23ace96-eb09-5bf7-b884-7ef8e8fc544d")
	config := game.ActionConfig{
		Name:        "Leaf Jab",
		Description: "+1 priority.",
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(50),
		Stat:        game.Ptr(game.StatAttack),
		Nature:      game.Ptr(game.NsWood),
		Cost:        game.Ptr(30),
		TargetCount: game.Ptr(1),
		Jutsu:       game.Taijutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	action := makeBasicAttack(ID, config)
	action.Priority = game.ActionPriorityP1
	return action
}
