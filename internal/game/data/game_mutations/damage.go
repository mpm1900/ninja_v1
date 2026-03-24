package mutations

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func ApplyDamage(g *game.Game, target game.ResolvedActor, damage int) int {
	clamped := Clamp(damage, 0, damage)
	g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
		a.Damage += clamped
		a.Alive = target.Stats[game.StatHP] > a.Damage
		return a
	})

	return clamped
}

func PureDamage(damage int) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			targets := g.GetTargets(context)
			for _, t := range targets {
				target := game.ResolveActor(t, g)
				ApplyDamage(&g, target, damage)

			}

			return g
		},
	}
}

func NewDamage(config game.ActionConfig) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			ok, s := g.GetSource(context)
			if !ok {
				return g
			}

			targets := g.GetTargets(context)
			source := game.ResolveActor(s, g)
			total := 0
			for _, t := range targets {
				target := game.ResolveActor(t, g)
				damages := game.GetDamage(
					source,
					[]game.ResolvedActor{target},
					*config.Stat,
					*config.Power,
					config.Nature,
				)

				for _, damage := range damages {
					total += ApplyDamage(&g, target, damage)
				}
			}

			if config.Recoil != nil && *config.Recoil > 0 && context.SourceActorID != nil {
				recoilContext := game.Context{
					ParentActorID:  context.ParentActorID,
					SourceActorID:  context.SourceActorID,
					SourcePlayerID: context.SourcePlayerID,
					// Set the source as the target
					TargetActorIDs: []uuid.UUID{*context.SourceActorID},
				}
				damageMut := PureDamage(int(*config.Recoil * float64(total)))
				damageTx := game.MakeTransaction(damageMut, recoilContext)
				g.JumpTransaction(damageTx)
			}

			return g
		},
	}

}

func MakeDamageTransactions(context game.Context, damages ...game.GameMutation) []game.GameTransaction {
	var transactions []game.GameTransaction
	for _, damage := range damages {
		transactions = append(
			transactions,
			game.MakeTransaction(
				damage,
				context,
			),
		)
	}

	return transactions
}
