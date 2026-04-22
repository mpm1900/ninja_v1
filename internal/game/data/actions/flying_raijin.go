package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var FlyingRaijin = MakeFlyingRaijin()

func MakeFlyingRaijin() game.Action {
	ID := uuid.MustParse("1a54031e-0ae6-49ed-b8b5-931c692417bf")
	config := game.ActionConfig{
		Name:        "Flying Raijin",
		Description: "+2 priority. High critical hit chance.",
		Nature:      game.Ptr(game.NsTai),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(80),
		Stat:        game.Ptr(game.StatAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(0),
		Cooldown:    game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(getCriticalStage(1)),
		CritMod:     1.5,
	}

	action := makeBasicAttack(ID, config)
	action.Priority = game.ActionPriorityP2
	return action
}
