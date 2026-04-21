package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var ChidoriStream = MakeChidoriStream()

func MakeChidoriStream() game.Action {
	ID := uuid.MustParse("8973e59d-dc43-4eac-99e2-68e6ca114aa9")

	config := game.ActionConfig{
		Name:        "Chidori Stream",
		Description: "Hits all enemy shinobi. 20% chance to paralyze targets.",
		Nature:      game.Ptr(game.NsLightning),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(80),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(0),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(0),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	action := makeBasicAttackWith(ID, config, func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction {
		targets := g.GetTargets(context)
		for _, target := range targets {
			transactions = append(transactions, chanceParalysis(config, target, 20)...)
		}

		return transactions
	}, nil)
	action.TargetPredicate = game.NoneFilter
	action.MapContext = func(g game.Game, context game.Context) game.Context {
		other_actors := g.GetActorsFilters(context, game.ComposeAF(game.ActiveFilter, game.OtherTeamFilter))
		for _, t := range other_actors {
			context.TargetPositionIDs = append(context.TargetPositionIDs, *t.PositionID)
		}
		return context
	}

	return action
}
