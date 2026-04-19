package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func SetWeather(gid uuid.UUID, weather game.GameWeather, name string) game.Modifier {
	return game.Modifier{
		ID:       gid,
		GroupID:  &gid,
		Name:     name,
		Show:     true,
		Duration: game.ModifierDurationInf,
		GameStateMutations: []game.GameStateMutation{
			game.MakeGameStateMutation(
				&gid,
				game.MutPriorityGameState0,
				game.GS_TrueFilter,
				func(g game.Game, gs game.GameState, context game.Context) game.GameState {
					gs.Weather = weather
					return gs
				},
			),
		},
		ActorMutations: []game.ActorMutation{},
		Triggers:       []game.Trigger{},
	}
}
