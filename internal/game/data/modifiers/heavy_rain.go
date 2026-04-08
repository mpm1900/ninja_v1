package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var heavyRainID = uuid.New()

var HeavyRain game.Modifier = game.Modifier{
	ID:       heavyRainID,
	GroupID:  &heavyRainID,
	Name:     "Heavy Rain",
	Show:     true,
	Duration: game.ModifierDurationInf,
	GameStateMutations: []game.GameStateMutation{
		game.MakeGameStateMutation(
			&heavyRainID,
			game.MutPriorityGameState0,
			game.GS_SourceIsActiveFilter,
			func(g game.Game, gs game.GameState, context game.Context) game.GameState {
				gs.Weather = game.GameWeatherRain
				return gs
			},
		),
	},
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&heavyRainID),
	},
	Triggers: []game.Trigger{},
}
