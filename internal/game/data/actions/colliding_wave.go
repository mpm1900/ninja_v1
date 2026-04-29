package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var CollidingWave = MakeCollidingWave()

func MakeCollidingWave() game.Action {
	ID := uuid.MustParse("74d5a7d7-cb62-58b4-9ace-e80bf7f0fd40")

	config := game.ActionConfig{
		Name:        "Colliding Wave",
		Description: "Hits all other active shinobi. Sets flooded terrain.",
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

	action := makeBasicAttackWith(
		ID,
		config,
		func(g game.Game, _, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}

			state, _ := g.GetState(context)
			if state.Terrain == game.GameTerrainFlooded {
				return transactions
			}

			filter := modifiers.FilterTerrain()
			transactions = append(transactions, filter)

			mod := modifiers.FloodedTerrain()
			mod.Duration = 4
			mut := mutations.AddModifiers(false, mod)
			transaction := game.MakeTransaction(mut, game.NewContext())
			transactions = append(transactions, transaction)

			return transactions
		},
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
