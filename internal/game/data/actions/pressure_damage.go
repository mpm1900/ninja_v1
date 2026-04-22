package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var PressureDamage = MakePressureDamage()

func MakePressureDamage() game.Action {
	ID := uuid.MustParse("22f9cca5-b709-444d-a1d1-72f3707b08cc")

	config := game.ActionConfig{
		Name:        "Pressure Damage",
		Description: "Hits all enemy shinobi. Grants the user Wind nature until end of turn.",
		Nature:      game.Ptr(game.NsWind),
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(75),
		Stat:        game.Ptr(game.StatChakraAttack),
		TargetCount: game.Ptr(0),
		Cost:        game.Ptr(30),
		Cooldown:    game.Ptr(1),
		Jutsu:       game.Ninjutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}

	action := makeBasicAttackWith(ID, config, nil, func(g game.Game, context game.Context, transactions []game.GameTransaction) []game.GameTransaction {
		add_mut := mutations.AddModifiers(false, modifiers.AddNature(game.NsWind, 0))
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
