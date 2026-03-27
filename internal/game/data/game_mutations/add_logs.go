package mutations

import (
	"ninja_v1/internal/game"
)

func AddLogs(logs ...game.GameLog) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			g.Log = append(g.Log, logs...)

			return g
		},
	}
}
