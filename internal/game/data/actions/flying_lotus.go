package actions

import (
	"fmt"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var flyingLotusID = uuid.MustParse("75d0a924-912a-4972-93fb-3c08e82bb1b3")
var FlyingLotus = MakeFlyingLotus()

func MakeFlyingLotus() game.Action {
	config := game.ActionConfig{
		Name:        "Flying Lotus",
		Description: "Must attack for 3 turns.",
		Nature:      game.Ptr(game.NsTai),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(100),
		Stat:        game.Ptr(game.StatAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(0),
		Jutsu:       game.Taijutsu,
		CritChance:  game.Ptr(getCriticalStage(0)),
		CritMod:     1.5,
	}
	action := makeBasicAttackWith(
		flyingLotusID,
		config,
		func(g game.Game, context game.Context, _ game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}

			key := "repeats"
			repeats, ok := context.GetMeta(key)
			if !ok {
				repeats = 0
			}

			fmt.Println(repeats)
			fmt.Printf("%+v", context)
			context.SetMeta(key, repeats+1)
			if repeats < 2 {
				recharge := mutations.QueueAction(flyingLotusID, context)
				transactions = append(transactions, game.MakeTransaction(recharge, context))
			}

			return transactions
		},
		nil,
	)
	return action

}
