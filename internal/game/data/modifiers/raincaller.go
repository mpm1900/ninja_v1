package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var heavyRainID = uuid.MustParse("b28933cd-ad6a-5b83-acff-4dd084dad6e5")
var raincallerID = uuid.MustParse("912e5e72-263e-5f6f-8f9f-32d1a746cc49")
var HeavyRain = SetWeather(heavyRainID, game.GameWeatherRain, "Heavy Rain")

var RaincallerTrigger game.Trigger = game.Trigger{
	ID:         uuid.MustParse("21121daa-3823-5613-ae0d-c26f7f97fece"),
	ModifierID: raincallerID,
	On:         game.OnActorEnter,
	Check:      game.Match__SourceActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}

			mod := HeavyRain
			mod.Duration = 4
			mutation := mutations.AddModifiers(false, mod)
			transaction := game.MakeTransaction(mutation, game.NewContext())
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var Raincaller game.Modifier = game.Modifier{
	ID:          raincallerID,
	GroupID:     &raincallerID,
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
