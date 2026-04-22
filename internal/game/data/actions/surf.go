package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Surf = MakeSurf()

func MakeSurf() game.Action {
	ID := uuid.MustParse("74d5a7d7-cb62-58b4-9ace-e80bf7f0fd40")

	config := game.ActionConfig{
		Name:        "Surf",
		Description: "Hits all other active shinobi.",
		Nature:      game.Ptr(game.NsWater),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(90),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(0),
		Cost:        game.Ptr(30),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(getCriticalStage(0)),
		CritMod:     1.5,
	}

	action := makeBasicAttack(ID, config)
	action.TargetPredicate = game.NoneFilter
	action.MapContext = func(g game.Game, context game.Context) game.Context {
		other_actors := g.GetActorsFilters(context, game.ComposeAF(game.ActiveFilter, game.OtherFilter))
		for _, t := range other_actors {
			context.TargetPositionIDs = append(context.TargetPositionIDs, *t.PositionID)
		}
		return context
	}

	return action
}
