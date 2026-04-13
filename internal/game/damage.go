package game

import (
	"fmt"
	"maps"
	"slices"
)

type DamageTerms struct {
	Attack   int
	Critical float64
	Defense  int
	Level    int
	Nature   float64
	Offset   int
	Other    float64
	Power    int
	Random   float64
	STAB     float64
	Targets  float64
}

type DamageConfig struct {
	Critical        float64
	Random          float64
	IgnoreModifiers bool
	Repeat          bool
	RepeatMax       int
}

func NewDamageConfig(critical float64, random float64) DamageConfig {
	return DamageConfig{
		Critical:        critical,
		Random:          random,
		IgnoreModifiers: critical > 1,
		Repeat:          false,
		RepeatMax:       0,
	}
}

func DamageEquation(terms DamageTerms) int {
	pow_ad := float64(terms.Power) * float64(terms.Attack) / float64(terms.Defense)
	level_mod := float64(2*terms.Level)/5 + 2
	base := (pow_ad*level_mod)/50 + 2
	raw := (base * terms.Critical * terms.Nature * terms.STAB * terms.Targets * terms.Random * terms.Other)
	fmt.Printf(
		"(((%d * %d / %d) * ((2 * %d / 5) + 2) / 50 + 2) * (%f * %f * %f * %f * %f)  = %f \n",
		terms.Power, terms.Attack, terms.Defense, terms.Level, terms.Nature, terms.STAB, terms.Targets, terms.Random, terms.Other, raw,
	)
	return Round(raw) + terms.Offset
}

func GetStabModifier(source ResolvedActor, nature *NatureSet) float64 {
	if nature == nil {
		return 1.00
	}

	natures := slices.Collect(maps.Keys(source.Natures))
	index := slices.IndexFunc(natures, func(n NatureSet) bool {
		return n == *nature
	})

	if index == -1 {
		return 1.00
	}

	return source.StabMultiplier
}

func HasDebuff(r ResolvedActor, stat AttackStat) bool {
	return r.DamageMultipliers[stat] < 1 ||
		r.Stages[ActorStat(stat)] < 0 ||
		r.Stages[StatAccuracy] < 0
}
func HasBuff(r ResolvedActor, attack AttackStat, defense DefenseStat) bool {
	return r.DamageReduction[attack] > 1 ||
		r.Stages[ActorStat(defense)] > 0 ||
		r.Stages[StatEvasion] > 0
}

func GetDamage(
	source ResolvedActor,
	targets []ResolvedActor,
	ignoreModifiers bool,
	targetsCount int,
	attack AttackStat,
	defense DefenseStat,
	power int,
	critical float64,
	nature *NatureSet,
	random float64,
) []int {
	damages := make([]int, len(targets))
	if power == 0 {
		return damages
	}

	a_base := float64(source.Stats[ActorStat(attack)])
	a_mod := 1.0
	attack_value := Round(a_base * a_mod)
	targets_mod := 1.0
	if targetsCount > 1 {
		targets_mod = 0.75
	}

	for i, target := range targets {
		d_base := float64(target.Stats[ActorStat(defense)])
		/**
		 * This piece is important. Critical hits ignore target stat changes.
		 */
		if critical > 1.0 {
			ignoreModifiers = true
		}
		if ignoreModifiers {
			if HasBuff(target, attack, defense) {
				d_base = float64(target.UnmodifiedStats[ActorStat(defense)])
			}
			if HasDebuff(source, attack) {
				a_base = float64(source.UnmodifiedStats[ActorStat(attack)])
			}
		}
		d_mod := 1.0
		defense_value := Round(d_base * d_mod)

		var natures []Nature
		if nature != nil {
			natures = NATURES[*nature]
		}
		nature_mod := ResolveNatures(natures, source.NatureDamage, target.NatureResistance, target.Natures)
		stab_mod := GetStabModifier(source, nature)

		damages[i] = DamageEquation(DamageTerms{
			Attack:   attack_value,
			Critical: critical,
			Defense:  defense_value,
			Level:    source.Level,
			Nature:   nature_mod,
			Offset:   0,
			Other:    source.DamageMultipliers[attack] * target.DamageReduction[attack],
			Power:    Round(float64(power) * source.PowerMultiplier),
			Random:   random,
			STAB:     stab_mod,
			Targets:  targets_mod,
		})
	}
	return damages
}
