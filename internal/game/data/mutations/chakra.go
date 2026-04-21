package mutations

import (
	"ninja_v1/internal/game"
)

func UseStaminaSource(amount int) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			if context.SourceActorID == nil {
				return g
			}

			g.UpdateActor(*context.SourceActorID, func(a game.Actor) game.Actor {
				a.StaminaDamage += amount
				return a
			})

			return g
		},
	}
}

func GainStaminaSource(amount int) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			if context.SourceActorID == nil {
				return g
			}

			g.UpdateActor(*context.SourceActorID, func(a game.Actor) game.Actor {
				a.StaminaDamage = max(a.StaminaDamage-amount, 0)
				return a
			})

			return g
		},
	}
}

func RecoverStaminaSource(ratio float64) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			if context.SourceActorID == nil {
				return g
			}

			g.UpdateActor(*context.SourceActorID, func(a game.Actor) game.Actor {
				a.RecoverStamina(g, ratio)
				return a
			})

			return g
		},
	}
}
