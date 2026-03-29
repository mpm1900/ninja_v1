package mutations

import (
	"fmt"
	"ninja_v1/internal/game"
)

func UseChakraSource(amount int) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			if context.SourceActorID == nil {
				return g
			}

			g.UpdateActor(*context.SourceActorID, func(a game.Actor) game.Actor {
				fmt.Printf("%d, %d\n", a.ChakraDamage, amount)
				a.ChakraDamage += amount
				return a
			})

			return g
		},
	}
}

func GainChakraSource(amount int) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			if context.SourceActorID == nil {
				return g
			}

			g.UpdateActor(*context.SourceActorID, func(a game.Actor) game.Actor {
				a.ChakraDamage = max(a.ChakraDamage - amount, 0)
				return a
			})

			return g
		},
	}
}

func RecoverChakraSource(ratio float64) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			if context.SourceActorID == nil {
				return g
			}

			g.UpdateActor(*context.SourceActorID, func(a game.Actor) game.Actor {
				a.RecoverChakra(g, ratio)
				return a
			})

			return g
		},
	}
}
