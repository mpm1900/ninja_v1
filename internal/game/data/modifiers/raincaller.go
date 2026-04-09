package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var heavyRainID = uuid.New()

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

var raincallerID = uuid.New()

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
