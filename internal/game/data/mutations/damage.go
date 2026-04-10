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
	for i, t := range targets {
		resolved[i] = t.Resolve(g)
	}
	return resolved
}

func getDefenseStat(a game.AttackStat) game.DefenseStat {
	if a == game.ChakraAttack {
		return game.ChakraDefense
	}
	return game.Defense
}

// returns if target is still alive after
func ApplyDamageWith(g *game.Game, source_ID *uuid.UUID, target game.ResolvedActor, damage int, updater func(game.Actor) game.Actor) bool {
	alive := target.Alive
	if !target.Alive {
		return alive
	}

	logCtx := game.NewContext()
	logCtx.ParentActorID = &target.ID
	logCtx.SourceActorID = &target.ID

	hp := target.Stats[game.StatHP]

	g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
		if a.Summon != nil && a.Summon.Proxy && a.Summon.Alive {
			summonHP := a.Summon.Stats[game.StatHP]
			a.Summon.Damage += damage
			a.Summon.Alive = summonHP > a.Summon.Damage
			g.PushLog(game.NewLogContext("| $source$'s summon took the attack.", logCtx))
		} else {
			a.Damage += clampDamage(damage)
			if source_ID != nil {
				a.LastRecievedDamage[*source_ID] = clampDamage(damage)
			}
			ratio := min(int(float64(damage)*100/float64(hp)), 100)
			if ratio > 0 {
				g.PushLog(game.NewLogContext(fmt.Sprintf("| $source$ lost %d%% HP.", ratio), logCtx))
			} else {
				g.PushLog(game.NewLogContext(fmt.Sprintf("| $source$ gained %d%% HP.", ratio*-1), logCtx))
			}

			if target.Immortal && hp <= a.Damage {
				g.PushLog(game.NewLogContext("| $source$'s survived a fatal attack.", logCtx))
				a.Damage = hp - 1
				g.On(game.OnImmortalSave, &logCtx)
			}

			a.Alive = hp > a.Damage
			alive = a.Alive
		}

		if updater == nil {
			return a
		}
		return updater(a)
	})

	return alive
}

// returns if target is still alive after
func ApplyDamage(g *game.Game, source_ID *uuid.UUID, target game.ResolvedActor, damage int) bool {
	return ApplyDamageWith(g, source_ID, target, damage, nil)
}

func PureDamageWith(damage int, trigger bool, updater func(game.Actor) game.Actor) game.GameMutation {
	return game.GameMutation{
		Filter: game.TargetsIsOneAlive,
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			targets := g.GetTargets(context)
			for _, t := range targets {
				target := t.Resolve(g)
				ApplyDamageWith(&g, context.SourceActorID, target, damage, updater)
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
				ApplyDamageWith(&g, context.SourceActorID, target, damage, updater)
			}
			return g
		},
	}
}

func RatioDamage(ratio float64) game.GameMutation {
	return RatioDamageWith(ratio, nil)
}

type damageHandler struct {
	action   game.ActionConfig
	config   game.DamageConfig
	context  game.Context
	source   game.ResolvedActor
	resolved []game.ResolvedActor
	defense  game.DefenseStat

	total  int
	totals []int

	repeats            int
	repeatTransactions []game.GameTransaction
	sideEffectTxs      []game.GameTransaction
}

func newDamageExecution(g game.Game, action game.ActionConfig, config game.DamageConfig, context game.Context, source game.ResolvedActor) *damageHandler {
	resolved := resolveTargets(g, context)
	return &damageHandler{
		action:             action,
		config:             config,
		context:            context,
		source:             source,
		resolved:           resolved,
		defense:            getDefenseStat(*action.Stat),
		totals:             make([]int, len(resolved)),
		repeatTransactions: make([]game.GameTransaction, 0),
		sideEffectTxs:      make([]game.GameTransaction, 0),
	}
}

func (e *damageHandler) run(g *game.Game) {
	for {
		missed := false

		for ti, target := range e.resolved {
			if e.resolveTargetHit(g, ti, target) {
				missed = true
			}
		}

		if !e.config.Repeat || missed {
			break
		}

		if e.config.RepeatMax < 0 || e.config.RepeatMax > e.repeats {
			e.repeats++
		} else {
			break
		}
	}

	e.buildSideEffects()
	e.commitTransactions(g)
}

func (e *damageHandler) resolveTargetHit(g *game.Game, targetIndex int, target game.ResolvedActor) bool {
	if target.Protected {
		g.PushLog(game.NewLogContext("| $source$ was protected.", e.context.WithSource(target.ID)))
		return false
	}

	result := game.MakeAccuracyCheck(g, e.action, e.source, target)
	if !result.Success {
		if !e.config.Repeat || e.repeats == 0 {
			g.PushLog(game.NewLog(fmt.Sprintf("%s missed!", e.action.Name)))
			g.PushLog(game.NewLog(fmt.Sprintf("roll = %d, acc = %d", result.Roll, result.Chance)))
		}
		return true
	}

	damages := game.GetDamage(
		e.source,
		[]game.ResolvedActor{target},
		e.config.IgnoreStages,
		len(e.resolved),
		*e.action.Stat,
		e.defense,
		*e.action.Power,
		e.config.Critical,
		e.action.Nature,
		e.config.Random,
	)

	e.logNatureEffectiveness(g, target)

	for _, damage := range damages {
		if e.config.Repeat {
			e.queueRepeatHit(target, damage)
		} else {
			e.applySingleHit(g, target, damage)
		}

		applied := clampDamage(damage)
		e.total += applied
		e.totals[targetIndex] += applied
	}

	return false
}
func (e *damageHandler) applySingleHit(g *game.Game, target game.ResolvedActor, damage int) {
	alive := ApplyDamage(g, &e.source.ID, target, damage)

	if damage > 0 {
		g.On(game.OnDamageRecieve, &e.context)
	}

	if e.config.Critical > 1.0 {
		g.PushLog(game.NewLog(fmt.Sprintf("Critical Hit! (x%f)", e.config.Critical)))
	}

	if !alive {
		deathContext := game.NewContext().WithSource(target.ID)
		killContext := game.NewContext().WithSource(e.source.ID).WithTargetIDs([]uuid.UUID{target.ID})
		g.On(game.OnDeath, &deathContext)
		g.On(game.OnKill, &killContext)
	}
}
func (e *damageHandler) queueRepeatHit(target game.ResolvedActor, damage int) {
	targetContext := e.context
	targetContext.TargetActorIDs = []uuid.UUID{target.ID}
	targetContext.TargetPositionIDs = []uuid.UUID{}

	repeatTx := game.MakeTransaction(PureDamage(damage, true), targetContext)

	log := game.NewLogContext(fmt.Sprintf("$action$ hit %d times.", e.repeats+1), e.context)
	logMux := game.AddLogs(log)
	logMux.Filter = game.TargetsIsOneAlive
	logTx := game.MakeTransaction(logMux, e.context)

	e.repeatTransactions = append(e.repeatTransactions, logTx, repeatTx)
}
func (e *damageHandler) logNatureEffectiveness(g *game.Game, target game.ResolvedActor) {
	var natures []game.Nature
	if e.action.Nature != nil {
		natures = game.NATURES[*e.action.Nature]
	}

	natureMod := game.ResolveNatures(natures, e.source.NatureDamage, target.NatureResistance, target.Natures)
	if natureMod >= 1.5 {
		g.PushLog(game.NewLog("| Super effective!"))
	}
	if natureMod < 0.75 {
		g.PushLog(game.NewLog("| Not very effective!"))
	}
}
func (e *damageHandler) buildSideEffects() {
	if e.total <= 0 || e.context.SourceActorID == nil {
		return
	}

	context := game.MakeContextForActor(e.source.Actor)
	if e.action.LifeSteal != nil && *e.action.LifeSteal > 0.0 {
		amount := game.Round(*e.action.LifeSteal * float64(e.total))
		healTx := game.MakeTransaction(PureHeal(amount), context)
		e.sideEffectTxs = append(e.sideEffectTxs, healTx)
	}

	if e.action.Recoil != nil && *e.action.Recoil > 0.0 {
		amount := game.Round(*e.action.Recoil * float64(e.total))
		recoilTx := game.MakeTransaction(PureDamage(amount, false), context)
		e.sideEffectTxs = append(e.sideEffectTxs, recoilTx)
	}

	for i, target := range e.resolved {
		if target.Reflect > 0.0 && *e.context.SourceActorID != target.ID {
			reflectDamage := int(target.Reflect * float64(e.totals[i]))
			reflectTx := game.MakeTransaction(PureDamage(reflectDamage, false), context)
			e.sideEffectTxs = append(e.sideEffectTxs, reflectTx)
		}
	}
}
func (e *damageHandler) commitTransactions(g *game.Game) {
	ordered := make([]game.GameTransaction, 0, len(e.repeatTransactions)+len(e.sideEffectTxs))
	ordered = append(ordered, e.repeatTransactions...)
	ordered = append(ordered, e.sideEffectTxs...)

	for i := len(ordered) - 1; i >= 0; i-- {
		g.JumpTransaction(ordered[i])
	}
}

func NewDamage(action game.ActionConfig, config game.DamageConfig) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			s, ok := g.GetSource(context)
			if !ok || action.Stat == nil || action.Power == nil {
				return g
			}

			source := s.Resolve(g)
			exec := newDamageExecution(g, action, config, context, source)
			exec.run(&g)

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
		return updater(a)
	})

	t, ok := g.GetActorByID(targetID)
	if !ok {
		return amount
	}

	target := t.Resolve(*g)
	hp := target.Stats[game.StatHP]
	logCtx := game.MakeContextForActor(target.Actor)
	ratio := int(float64(amount) * 100 / float64(hp))
	g.PushLog(game.NewLogContext(fmt.Sprintf("| $source$ gained %d%% HP.", ratio), logCtx))

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
			for _, target := range resolveTargets(g, context) {
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

			source := s.Resolve(g)
			for _, target := range resolveTargets(g, context) {
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
