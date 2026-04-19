package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Rasenshuriken = MakeRasenshuriken()

func MakeRasenshuriken() game.Action {
	ID := uuid.MustParse("6b3df363-7052-47fc-af99-7e8eafdc9ee2")
	config := game.ActionConfig{
		Name:        "Rasenshuriken",
		Description: "User's Chakra Attack is lowered by 2 stages.",
		Nature:      game.Ptr(game.NsWind),
		Accuracy:    game.Ptr(90),
		Power:       game.Ptr(130),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(1),
		Cost:        game.Ptr(50),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	return makeBasicAttackWith(ID, config, func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction {
		source, ok := g.GetSource(context)
		if !ok {
			return transactions
		}

		mutation := mutations.AddModifiers(false, modifiers.ChakraAttackDown2Source)
		transaction := game.MakeTransaction(mutation, game.MakeContextForActor(source))
		transactions = append(transactions, transaction)

		return transactions
	}, nil)
}
