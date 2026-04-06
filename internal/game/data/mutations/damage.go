package mutations

import (
	"fmt"
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

	log_ctx := game.NewContext()
	log_ctx.ParentActorID = &target.ID
	log_ctx.SourceActorID = &target.ID

	hp := target.Stats[game.StatHP]

	g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
		if a.Summon != nil && a.Summon.Alive {
			summon_hp := a.Summon.Stats[game.StatHP]
			a.Summon.Damage += damage
			a.Summon.Alive = summon_hp > a.Summon.Damage
			g.PushLog(game.NewLogContext(">>> $source$'s substitute took the attack.", log_ctx))
		} else {
			a.Damage += damage
			a.Alive = hp > a.Damage
			ratio := int(float64(damage) * 100 / float64(hp))
			if ratio > 0 {
				g.PushLog(game.NewLogContext(fmt.Sprintf(">>> $source$ lost %d%% HP.", ratio), log_ctx))
			} else {
				g.PushLog(game.NewLogContext(fmt.Sprintf(">>> $source$ gained %d%% HP.", ratio*-1), log_ctx))
			}
		}

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
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			targets := g.GetTargets(context)
			for _, t := range targets {
				target := t.Resolve(g)
				ApplyDamageWith(&g, target, damage, updater)
				if trigger && damage > 0 {
					g.On(game.OnDamageRecieve, &context)
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
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			targets := g.GetTargets(context)
			for _, t := range targets {
				target := t.Resolve(g)
				damage := game.Round(float64(target.Stats[game.StatHP]) * ratio)
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
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
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

					result := game.MakeAccuracyCheck(&g, action, source, target)
					if !result.Success {
						if !config.Repeat || repeats == 0 {
							g.PushLog(game.NewLog(fmt.Sprintf("%s missed!", action.Name)))
							g.PushLog(game.NewLog(fmt.Sprintf("roll = %d, acc = %d", result.Roll, result.Chance)))
						}
						missed = true
						continue
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
								g.On(game.OnDamageRecieve, &context)
							}
							if config.Critical > 1.0 {
								g.PushLog(game.NewLog(fmt.Sprintf("Critical Hit! (x%f)", config.Critical)))
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
				amount := game.Round(*action.LifeSteal * float64(total))
				healMut := PureHeal(amount)
				damageTx := game.MakeTransaction(healMut, selfContext)
				sideEffectTransactions = append(sideEffectTransactions, damageTx)
			}

			if total > 0 && action.Recoil != nil && *action.Recoil > 0.0 && context.SourceActorID != nil {
				amount := game.Round(*action.Recoil * float64(total))
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

	t, ok := g.GetActorByID(targetID)
	if !ok {
		return amount
	}

	target := t.Resolve(*g)
	hp := target.Stats[game.StatHP]
	log_ctx := game.NewContext()
	log_ctx.ParentActorID = &targetID
	log_ctx.SourceActorID = &targetID
	ratio := int(float64(amount) * 100 / float64(hp))
	g.PushLog(game.NewLogContext(fmt.Sprintf(">>> $source$ gained %d%% HP.", ratio), log_ctx))

	return amount
}
func ApplyHealRaw(g *game.Game, targetID uuid.UUID, amount int) int {
	return ApplyHealRawWith(g, targetID, amount, nil)
}
func ApplyHealRatioWith(g *game.Game, target game.ResolvedActor, ratio float64, updater func(game.Actor) game.Actor) int {
	amount := game.Round(float64(target.Stats[game.StatHP]) * ratio)
	return ApplyHealRawWith(g, target.ID, amount, updater)
}
func ApplyHealRatio(g *game.Game, target game.ResolvedActor, ratio float64) int {
	return ApplyHealRatioWith(g, target, ratio, nil)
}

func RatioHeal(ratio float64) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
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
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
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
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
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
				result := game.MakeAccuracyCheck(&g, action, source, target)
				if !result.Success {
					g.PushLog(game.NewLog(fmt.Sprintf("%s missed!", action.Name)))
					g.PushLog(game.NewLog(fmt.Sprintf("roll = %d, acc = %d", result.Roll, result.Chance)))
					continue
				}

				ApplyHealRatio(&g, target, ratio)

			}

			return g
		},
	}
}
