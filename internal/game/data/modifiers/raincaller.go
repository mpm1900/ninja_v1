package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var heavyRainID = uuid.MustParse("b28933cd-ad6a-5b83-acff-4dd084dad6e5")

func HeavyRain() game.Modifier {
	return game.Modifier{
		ID:       heavyRainID,
		GroupID:  &heavyRainID,
		Name:     "Heavy Rain",
		Show:     true,
		Duration: game.ModifierDurationInf,
		GameStateMutations: []game.GameStateMutation{
			game.MakeGameStateMutation(
				&heavyRainID,
				game.MutPriorityGameState0,
				game.GS_TrueFilter,
				func(g game.Game, gs game.GameState, context game.Context) game.GameState {
					gs.Weather = game.GameWeatherRain
					return gs
				},
			),
		},
		ActorMutations: []game.ActorMutation{},
		Triggers:       []game.Trigger{},
	}
}

var raincallerID = uuid.MustParse("912e5e72-263e-5f6f-8f9f-32d1a746cc49")

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

			mod := HeavyRain()
			mod.Duration = 4
			mutation := mutations.AddModifiers(false, false, mod)
			transaction := game.MakeTransaction(mutation, game.NewContext())
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var Raincaller game.Modifier = game.Modifier{
	ID:       raincallerID,
	GroupID:  &raincallerID,
	Name:     "Raincaller",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&raincallerID),
	},
	Triggers: []game.Trigger{
		RaincallerTrigger,
	},
}
