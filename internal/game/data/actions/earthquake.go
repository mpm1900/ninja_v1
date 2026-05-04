package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Earthquake = MakeEarthquake()

func MakeEarthquake() game.Action {
	ID := uuid.MustParse("80e197af-299c-448d-a431-12473fa13866")

	config := game.ActionConfig{
		Name:        "Earthquake",
		Description: "Hits all other active shinobi.",
		Nature:      game.Ptr(game.NsEarth),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(100),
		Stat:        game.Ptr(game.StatAttack),
		TargetCount: game.Ptr(0),
		Cost:        game.Ptr(30),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(getCriticalStage(0)),
		CritMod:     1.5,
	}

	action := makeBasicAttackWith(
		ID,
		config,
		nil,
		nil,
	)
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
