package mutations

import (
	"ninja_v1/internal/game"
)

func AddModifiers(modifiers []game.Modifier) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			for _, modifier := range modifiers {
				g.Modifiers = append(g.Modifiers, game.MakeTransaction(modifier, context))
			}

			return g
		},
	}
}
