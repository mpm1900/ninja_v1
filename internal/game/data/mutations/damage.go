package mutations

import (
	"fmt"
	"math"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func clampDamage(damage int) int {
	if damage < 0 {
		return 0
	}
	return damage
}
func resolveTargets(g game.Game, context game.Context) []game.ResolvedActor {
	targets := g.GetTargets(context)
	resolved := make([]game.ResolvedActor, len(targets))
	for t_index, t := range targets {
		resolved[t_index] = t.Resolve(g)
	}

	return resolved
}

func ApplyDamageWith(g *game.Game, target game.ResolvedActor, damage int, updater func(game.Actor) game.Actor) int {
	if !target.Alive {
		return 0
	}

	clamped := clampDamage(damage)
	hp := target.Stats[game.StatHP]

	g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
		a.Damage += clamped
		a.Alive = hp > a.Damage
		if updater == nil {
			return a
		}

		u := updater
		return u(a)
	})

	return clamped
}
func ApplyDamage(g *game.Game, target game.ResolvedActor, damage int) int {
	return ApplyDamageWith(g, target, damage, nil)
}

func PureDamageWith(damage int, updater func(game.Actor) game.Actor) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			targets := g.GetTargets(context)
			for _, t := range targets {
				target := t.Resolve(g)
				ApplyDamageWith(&g, target, damage, updater)
			}

			return g
		},
	}
}
func PureDamage(damage int) game.GameMutation {
	return PureDamageWith(damage, nil)
}

func RatioDamageWith(ratio float64, updater func(game.Actor) game.Actor) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			targets := g.GetTargets(context)
			for _, t := range targets {
				target := t.Resolve(g)
				damage := int(math.Floor(float64(target.Stats[game.StatHP]) * ratio))
				ApplyDamageWith(&g, target, damage, updater)
			}

			return g
		},
	}
}
func RatioDamage(ratio float64) game.GameMutation {
	return RatioDamageWith(ratio, nil)
}

func NewDamage(action game.ActionConfig, config game.DamageConfig) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			ok, s := g.GetSource(context)
			if !ok {
				return g
			}

			source := s.Resolve(g)
			total := 0
			resolved := resolveTargets(g, context)
			totals := make([]int, len(resolved))
			for t_index, target := range resolved {
				if target.Protected {
					g.PushLog(fmt.Sprintf("%s was protected.", target.Name))
					continue
				}

				base_accuracy := game.GetAccuracy(g, source, target)

				if action.Accuracy != nil {
					accuracy := int(math.Floor(base_accuracy * float64(*action.Accuracy)))
					roll := game.MakeActionRoll()
					if roll > accuracy {
						g.PushLog(fmt.Sprintf("%s's %s missed!", source.Name, action.Name))
						g.PushLog(fmt.Sprintf("roll = %d, acc = %d", roll, accuracy))
						continue
					}
				}

				defense := game.Defense
				if *action.Stat == game.ChakraAttack {
					defense = game.ChakraDefense
				}

				damages := game.GetDamage(
					source,
					[]game.ResolvedActor{target},
					*action.Stat,
					defense,
					*action.Power,
					config.Critical,
					action.Nature,
					config.Random,
				)

				for _, damage := range damages {
					g.On(game.OnDamageRecieve, context)
					applied := ApplyDamage(&g, target, damage)
					total += applied
					totals[t_index] += applied
				}
			}

			/**
			 * LIFE STEAL
			 */
			if total > 0 && action.LifeSteal != nil && *action.LifeSteal > 0.0 && context.SourceActorID != nil {
				lifeStealContext := game.Context{
					ParentActorID:  context.SourceActorID,
					SourceActorID:  context.SourceActorID,
					SourcePlayerID: context.SourcePlayerID,
					// Set the source as the target
					TargetActorIDs: []uuid.UUID{*context.SourceActorID},
				}

				amount := int(math.Floor(*action.LifeSteal * float64(total)))
				healMut := PureHeal(amount)
				damageTx := game.MakeTransaction(healMut, lifeStealContext)
				g.JumpTransaction(damageTx)
			}

			/**
			 * RECOIL DAMAGE
			 */
			if total > 0 && action.Recoil != nil && *action.Recoil > 0.0 && context.SourceActorID != nil {
				recoilContext := game.Context{
					ParentActorID:  context.SourceActorID,
					SourceActorID:  context.SourceActorID,
					SourcePlayerID: context.SourcePlayerID,
					// Set the source as the target
					TargetActorIDs: []uuid.UUID{*context.SourceActorID},
				}
				amount := int(math.Floor(*action.Recoil * float64(total)))
				damageMut := PureDamage(amount)
				damageTx := game.MakeTransaction(damageMut, recoilContext)
				g.JumpTransaction(damageTx)
			}

			/*
			 * REFLECT DAMAGE
			 */
			if total > 0 && context.SourceActorID != nil {
				for t_index, target := range resolved {
					if target.Reflect > 0.0 && *context.SourceActorID != target.ID {
						reflectContext := game.Context{
							ParentActorID:  context.SourceActorID,
							SourceActorID:  context.SourceActorID,
							SourcePlayerID: context.SourcePlayerID,
							// Set the source as the target
							TargetActorIDs: []uuid.UUID{*context.SourceActorID},
						}
						damageMut := PureDamage(int(target.Reflect * float64(totals[t_index])))
						damageTx := game.MakeTransaction(damageMut, reflectContext)
						g.JumpTransaction(damageTx)
					}
				}
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

func ApplyHealRawWith(g *game.Game, targetID uuid.UUID, amount int, updater func(game.Actor) game.Actor) int {
	g.UpdateActor(targetID, func(a game.Actor) game.Actor {
		if !a.Alive {
			amount = 0
			return a
		}

		healed := min(amount, a.Damage)
		a.Damage -= healed
		amount = healed

		if updater == nil {
			return a
		}

		u := updater
		return u(a)
	})

	return amount
}
func ApplyHealRaw(g *game.Game, targetID uuid.UUID, amount int) int {
	return ApplyHealRawWith(g, targetID, amount, nil)
}
func ApplyHealRatioWith(g *game.Game, target game.ResolvedActor, ratio float64, updater func(game.Actor) game.Actor) int {
	amount := int(math.Floor(float64(target.Stats[game.StatHP]) * ratio))
	return ApplyHealRawWith(g, target.ID, amount, updater)
}
func ApplyHealRatio(g *game.Game, target game.ResolvedActor, ratio float64) int {
	return ApplyHealRatioWith(g, target, ratio, nil)
}

func RatioHeal(ratio float64) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			resolved := resolveTargets(g, context)
			for _, target := range resolved {
				ApplyHealRatio(&g, target, ratio)
			}
			return g
		},
	}
}

func PureHeal(amount int) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			targets := g.GetTargets(context)
			for _, target := range targets {
				ApplyHealRaw(&g, target.ID, amount)
			}
			return g
		},
	}
}

func NewHeal(action game.ActionConfig, ratio float64) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			ok, s := g.GetSource(context)
			if !ok {
				return g
			}

			targets := g.GetTargets(context)
			source := s.Resolve(g)
			resolved := make([]game.ResolvedActor, len(targets))
			for t_index, t := range targets {
				resolved[t_index] = t.Resolve(g)
			}

			for _, target := range resolved {
				base_accuracy := game.GetAccuracy(g, source, target)

				if action.Accuracy != nil {
					accuracy := int(math.Floor(base_accuracy * float64(*action.Accuracy)))
					roll := game.MakeActionRoll()
					if roll > accuracy {
						g.PushLog(fmt.Sprintf("%s's %s missed!", source.Name, action.Name))
						g.PushLog(fmt.Sprintf("roll = %d, acc = %d", roll, accuracy))
						continue
					}
				}

				ApplyHealRatio(&g, target, ratio)

			}
			return g
		},
	}
}
