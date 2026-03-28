package mutations

import "ninja_v1/internal/game"

func UseChakra(amount int) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			if context.SourceActorID == nil {
				return g
			}

			g.UpdateActor(*context.SourceActorID, func(a game.Actor) game.Actor {
				a.ChakraDamage += amount
				return a
			})

			return g
		},
	}
}
