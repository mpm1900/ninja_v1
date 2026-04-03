package game

import (
	"fmt"
	"maps"
	"math"
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
	Critical  float64
	Random    float64
	Repeat    bool
	RepeatMax int
}

func NewDamageConfig() DamageConfig {
	return DamageConfig{
		Critical:  1.0,
		Random:    1.0,
		Repeat:    false,
		RepeatMax: 0,
	}
}

func DamageEquation(terms DamageTerms) int {
	pow_ad := float64(terms.Power) * float64(terms.Attack) / float64(terms.Defense)
	level_mod := float64(2*terms.Level)/5 + 2
	base := (pow_ad*level_mod)/50 + 2
	raw := (base * terms.Critical * terms.Nature * terms.STAB * terms.Targets * terms.Random * terms.Other)
	fmt.Printf(
		"(((%d * %d / %d) * ((2 * %d / 5) + 2) / 50 + 2) * (%f * %f * %f * %f)  = %f \n",
		terms.Power, terms.Attack, terms.Defense, terms.Level, terms.Nature, terms.STAB, terms.Targets, terms.Random, raw,
	)
	return int(math.Floor(raw)) + terms.Offset
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

	return 1.5
}

func GetDamage(
	source ResolvedActor,
	targets []ResolvedActor,
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
	attack_value := int(math.Floor(a_base * a_mod))
	targets_mod := 1.0
	if targetsCount > 1 {
		targets_mod = 0.75
	}

	for i, target := range targets {
		d_base := float64(target.Stats[ActorStat(defense)])
		d_mod := 1.0
		defense_value := int(math.Floor(d_base * d_mod))

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
			Other:    1.00,
			Power:    power,
			Random:   random,
			STAB:     stab_mod,
			Targets:  targets_mod,
		})
	}
	return damages
}
