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
		a.State.Damage += clamped
		a.State.Alive = target.Stats[game.StatHP] > a.State.Damage
		return a
	})

	return clamped
}

func PureDamage(damage int) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			_, targets := g.GetTargets(context)
			for _, t := range targets {
				target := t.Resolve(g)
				ApplyDamage(&g, target, damage)
			}

			return g
		},
	}
}

func NewDamage(action game.ActionConfig, config game.DamageConfig) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			ok, s := g.GetSource(context)
			if !ok {
				return g
			}

			_, targets := g.GetTargets(context)
			source := s.Resolve(g)
			total := 0
			for _, t := range targets {
				target := t.Resolve(g)
				damages := game.GetDamage(
					source,
					[]game.ResolvedActor{target},
					*action.Stat,
					*action.Power,
					config.Critical,
					action.Nature,
					config.Random,
				)

				for _, damage := range damages {
					g.On(game.OnDamageRecieve, context)
					total += ApplyDamage(&g, target, damage)
				}
			}

			if total > 0 && action.Recoil != nil && *action.Recoil > 0 && context.SourceActorID != nil {
				recoilContext := game.Context{
					ParentActorID:  context.SourceActorID,
					SourceActorID:  context.SourceActorID,
					SourcePlayerID: context.SourcePlayerID,
					// Set the source as the target
					TargetActorIDs: []uuid.UUID{*context.SourceActorID},
				}
				damageMut := PureDamage(int(*action.Recoil * float64(total)))
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
