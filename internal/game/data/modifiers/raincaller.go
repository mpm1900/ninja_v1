package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var raincallerID = uuid.MustParse("912e5e72-263e-5f6f-8f9f-32d1a746cc49")

var RaincallerTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: raincallerID,
	On:         game.OnActorEnter,
	Check:      game.Match__SourceActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}

			state, _ := g.GetState(context)
			if state.Weather == game.GameWeatherRain {
				return transactions
			}

			filter := FilterWeather()
			transactions = append(transactions, filter)

			mod := RainWeather()
			mod.Duration = 4
			mut := mutations.AddModifiers(false, mod)
			transaction := game.MakeTransaction(mut, game.NewContext())
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var Raincaller game.Modifier = game.Modifier{
	ID:          raincallerID,
	GroupID:     &raincallerID,
	Icon:        "raincaller",
	Name:        "Raincaller",
	Description: "On enter: start rain.",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&raincallerID),
	},
	Triggers: []game.Trigger{
		RaincallerTrigger,
	},
}
