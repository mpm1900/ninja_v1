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

func ApplyDamageWith(g *game.Game, target game.ResolvedActor, damage int, updater func(game.Actor) game.Actor) {
	if !target.Alive {
		return
	}

	hp := target.Stats[game.StatHP]

	g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
		a.Damage += damage // clamped
		a.Alive = hp > a.Damage
		if updater == nil {
			return a
		}

		u := updater
		return u(a)
	})
}
func ApplyDamage(g *game.Game, target game.ResolvedActor, damage int) {
	ApplyDamageWith(g, target, damage, nil)
}

func PureDamageWith(damage int, trigger bool, updater func(game.Actor) game.Actor) game.GameMutation {
	return game.GameMutation{
		Filter: game.TargetsIsOneAlive,
		Delta: func(g game.Game, context game.Context) game.Game {
			targets := g.GetTargets(context)
			for _, t := range targets {
				target := t.Resolve(g)
				ApplyDamageWith(&g, target, damage, updater)
				if trigger {
					g.On(game.OnDamageRecieve, context)
				}
			}

			return g
		},
	}
}
func PureDamage(damage int, trigger bool) game.GameMutation {
	return PureDamageWith(damage, trigger, nil)
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
			s, ok := g.GetSource(context)
			if !ok {
				return g
			}

			if action.Stat == nil || action.Power == nil {
				return g
			}

			selfContext := game.Context{
				ParentActorID:  context.ParentActorID,
				SourceActorID:  context.SourceActorID,
				SourcePlayerID: context.SourcePlayerID,
				TargetActorIDs: []uuid.UUID{*context.SourceActorID},
			}

			total := 0
			repeats := 0
			source := s.Resolve(g)
			resolved := resolveTargets(g, context)
			totals := make([]int, len(resolved))
			repeatTransactions := make([]game.GameTransaction, 0)

			defense := game.Defense
			if *action.Stat == game.ChakraAttack {
				defense = game.ChakraDefense
			}

			for {
				missed := false

				for ti, target := range resolved {
					if target.Protected {
						g.PushLog(game.NewLog(fmt.Sprintf("%s was protected.", target.Name)))
						continue
					}

					baseAccuracy := game.GetAccuracy(g, source, target)
					if action.Accuracy != nil {
						accuracy := int(math.Floor(baseAccuracy * float64(*action.Accuracy)))
						roll := game.MakeActionRoll()
						if roll > accuracy {
							if !config.Repeat || repeats == 0 {
								g.PushLog(game.NewLog(fmt.Sprintf("%s missed!", action.Name)))
								g.PushLog(game.NewLog(fmt.Sprintf("roll = %d, acc = %d", roll, accuracy)))
							}
							missed = true
							continue
						}
					}

					damages := game.GetDamage(
						source,
						[]game.ResolvedActor{target},
						len(resolved),
						*action.Stat,
						defense,
						*action.Power,
						config.Critical,
						action.Nature,
						config.Random,
					)

					for _, damage := range damages {

						if !config.Repeat {
							ApplyDamage(&g, target, damage)
							if damage > 0 {
								g.On(game.OnDamageRecieve, context)
							}

							total += clampDamage(damage)
							totals[ti] += clampDamage(damage)
							continue
						}

						targetContext := context
						targetContext.TargetActorIDs = []uuid.UUID{target.ID}
						targetContext.TargetPositionIDs = []uuid.UUID{}

						repeatTx := game.MakeTransaction(PureDamage(damage, true), targetContext)
						log := game.NewLogContext(fmt.Sprintf("$action$ hit %d times.", repeats+1), context)
						logMux := game.AddLogs(log)
						logMux.Filter = game.TargetsIsOneAlive
						logTx := game.MakeTransaction(logMux, context)
						repeatTransactions = append(repeatTransactions, logTx, repeatTx)

						applied := clampDamage(damage)
						total += applied
						totals[ti] += applied
					}
				}

				if !config.Repeat || missed {
					break
				}

				if config.RepeatMax < 0 || config.RepeatMax > repeats {
					repeats++
				} else {
					break
				}
			}

			sideEffectTransactions := make([]game.GameTransaction, 0)

			if total > 0 && action.LifeSteal != nil && *action.LifeSteal > 0.0 && context.SourceActorID != nil {
				amount := int(math.Floor(*action.LifeSteal * float64(total)))
				healMut := PureHeal(amount)
				damageTx := game.MakeTransaction(healMut, selfContext)
				sideEffectTransactions = append(sideEffectTransactions, damageTx)
			}

			if total > 0 && action.Recoil != nil && *action.Recoil > 0.0 && context.SourceActorID != nil {
				amount := int(math.Floor(*action.Recoil * float64(total)))
				damageMut := PureDamage(amount, false)
				damageTx := game.MakeTransaction(damageMut, selfContext)
				sideEffectTransactions = append(sideEffectTransactions, damageTx)
			}

			if total > 0 && context.SourceActorID != nil {
				for ti, target := range resolved {
					if target.Reflect > 0.0 && *context.SourceActorID != target.ID {
						damageMut := PureDamage(int(target.Reflect*float64(totals[ti])), false)
						damageTx := game.MakeTransaction(damageMut, selfContext)
						sideEffectTransactions = append(sideEffectTransactions, damageTx)
					}
				}
			}

			orderedTransactions := make([]game.GameTransaction, 0, len(repeatTransactions)+len(sideEffectTransactions))
			orderedTransactions = append(orderedTransactions, repeatTransactions...)
			orderedTransactions = append(orderedTransactions, sideEffectTransactions...)
			for i := len(orderedTransactions) - 1; i >= 0; i-- {
				g.JumpTransaction(orderedTransactions[i])
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
			s, ok := g.GetSource(context)
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
						g.PushLog(game.NewLog(fmt.Sprintf("%s missed!", action.Name)))
						g.PushLog(game.NewLog(fmt.Sprintf("roll = %d, acc = %d", roll, accuracy)))
						continue
					}
				}

				ApplyHealRatio(&g, target, ratio)

			}
			return g
		},
	}
}
