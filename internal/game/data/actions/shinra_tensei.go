package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var ShinraTensei = MakeShinraTensei()

func MakeShinraTensei() game.Action {
	ID := uuid.MustParse("fc9fe2ef-32e8-4e20-821e-b98b0a8fb4b7")

	config := game.ActionConfig{
		Name:        "Shinra Tensei",
		Description: "Hits all enemy shinobi.",
		Nature:      game.Ptr(game.NsYinYang),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(100),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(0),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(1),
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
		other_actors := g.GetActorsFilters(context, game.ComposeAF(game.ActiveFilter, game.OtherTeamFilter))
		for _, t := range other_actors {
			context.TargetPositionIDs = append(context.TargetPositionIDs, *t.PositionID)
		}
		return context
	}

	return action
}
