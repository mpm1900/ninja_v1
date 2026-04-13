package mutations

import (
	"math/rand"
	"ninja_v1/internal/game"
)

var Sleep = game.GameMutation{
	Delta: func(p, g game.Game, context game.Context) game.Game {
		targets := g.GetTargets(context)
		for _, target := range targets {
			g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
				if a.Statused {
					return a
				}

				a.SleepCounter = rand.Intn(3) + 1
				a.Sleeping = true
				a.Statused = true

				return a
			})
		}

		return g
	},
}

var Burn = game.GameMutation{
	Delta: func(p, g game.Game, context game.Context) game.Game {
		targets := g.GetTargets(context)
		for _, target := range targets {
			g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
				if a.Statused {
					return a
				}

				a.Burned = true
				a.Statused = true

				return a
			})
		}

		return g
	},
}

var Paralyze = game.GameMutation{
	Delta: func(p, g game.Game, context game.Context) game.Game {
		targets := g.GetTargets(context)
		for _, target := range targets {
			g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
				if a.Statused {
					return a
				}

				a.Paralyzed = true
				a.Statused = true

				return a
			})
		}

		return g
	},
}
