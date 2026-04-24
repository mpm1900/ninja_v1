package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/modifiers"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var CreateRain = MakeCreateRain()

func MakeCreateRain() game.Action {
	config := game.ActionConfig{
		Name:        "Create Rain",
		Nature:      game.Ptr(game.NsWater),
		Jutsu:       game.Ninjutsu,
		Description: "Creates rain.",
	}
	return game.Action{
		ID:              uuid.MustParse("2e05db18-ac05-48fd-9cf4-b50fc6e5dbc3"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				state, _ := g.GetState(context)
				if state.Weather == game.GameWeatherRain {
					return transactions
				}

				filter := modifiers.FilterWeather()
				transactions = append(transactions, filter)

				mod := modifiers.RainWeather()
				mod.Duration = 4
				mut := mutations.AddModifiers(false, mod)
				transaction := game.MakeTransaction(mut, game.NewContext())
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
