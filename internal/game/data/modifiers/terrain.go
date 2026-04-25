package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func FilterTerrain() game.GameTransaction {
	filter := game.NewGameMutation()
	filter.Delta = func(p, g game.Game, context game.Context) game.Game {
		g.FilterModifiers(func(mod game.Transaction[game.Modifier]) bool {
			return !mod.Mutation.Terrain
		})
		return g
	}
	return game.MakeTransaction(filter, game.NewContext())
}

func SetTerrain(gid uuid.UUID, terrain game.GameTerrain) game.Modifier {
	return game.Modifier{
		ID:       gid,
		GroupID:  &gid,
		Show:     true,
		Terrain:  true,
		Duration: game.ModifierDurationInf,
		GameStateMutations: []game.GameStateMutation{
			game.MakeGameStateMutation(
				&gid,
				game.MutPriorityGameState0,
				game.GS_TrueFilter,
				func(g game.Game, gs game.GameState, context game.Context) game.GameState {
					gs.Terrain = terrain
					return gs
				},
			),
		},
		ActorMutations: []game.ActorMutation{},
		Triggers:       []game.Trigger{},
	}
}

var floodedTerrainID = uuid.MustParse("f1784ee6-ba9a-4672-8eb5-e44b62021fec")

func FloodedTerrain() game.Modifier {
	mod := SetTerrain(floodedTerrainID, game.GameTerrainFlooded)
	mod.Name = "Flooded Terrain"
	mod.Icon = "flooded"
	mod.Description = "<< TODO >>"
	mod.ActorMutations = []game.ActorMutation{
		game.MakeActorMutation(
			&floodedTerrainID,
			game.MutPriorityDefault,
			game.ComposeAF(game.ActiveFilter, game.GameHasTerrain(game.GameTerrainFlooded)),
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				state, _ := g.GetState(context)
				if state.Terrain != game.GameTerrainFlooded {
					return actor
				}

				// TODO
				return actor
			},
		),
	}
	return mod
}
