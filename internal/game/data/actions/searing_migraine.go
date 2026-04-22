package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var SearingMigraine = MakeSearingMigraine()

func MakeSearingMigraine() game.Action {
	ID := uuid.MustParse("dc6edab6-535f-508f-b791-e197283eae86")

	config := game.ActionConfig{
		Name:        "Searing Migraine",
		Description: "Hits all enemy shinobi. Grants the user Fire nature until end of turn.",
		Nature:      game.Ptr(game.NsFire),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(75),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(0),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(getCriticalStage(0)),
		CritMod:     1.5,
	}

	action := makeBasicAttackWith(ID, config, nil, func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction {
		add_mut := mutations.AddModifiers(false, modifiers.AddNature(game.NsFire, 0))
		add_tx := game.MakeTransaction(add_mut, context)
		transactions = append(transactions, add_tx)
		return transactions
	})
	action.TargetPredicate = game.NoneFilter
	action.MapContext = func(g game.Game, context game.Context) game.Context {
		other_team_actors := g.GetActorsFilters(context, game.ComposeAF(game.ActiveFilter, game.OtherTeamFilter))
		for _, t := range other_team_actors {
			context.TargetPositionIDs = append(context.TargetPositionIDs, *t.PositionID)
		}
		return context
	}
	return action
}
