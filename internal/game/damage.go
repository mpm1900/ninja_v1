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
}

type DamageConfig struct {
	Critical float64
	Random   float64
}

func NewDamageConfig() DamageConfig {
	return DamageConfig{
		Critical: 1,
		Random:   1,
	}
}

func DamageEquation(terms DamageTerms) int {
	pow_ad := float64(terms.Power) * float64(terms.Attack) / float64(terms.Defense)
	level_mod := float64(2*terms.Level)/5 + 2
	base := (pow_ad*level_mod)/50 + 2
	raw := (base * terms.Critical * terms.Nature * terms.STAB * terms.Random * terms.Other)
	fmt.Printf(
		"(((%d * %d / %d) * ((2 * %d / 5) + 2) / 50 + 2) * (%f) * (%f) = %f \n",
		terms.Power, terms.Attack, terms.Defense, terms.Level, terms.Nature, terms.STAB, raw,
	)
	return int(math.Floor(raw)) + terms.Offset
}

func GetNaturesModifier(source, target ResolvedActor, moveNatures []Nature) float64 {
	targetNatures := make(map[Nature]struct{})
	for _, group := range target.Natures {
		for _, nature := range group {
			targetNatures[nature] = struct{}{}
		}
	}

	effectiveness := 1.0
	for _, moveNature := range moveNatures {
		for targetNature := range targetNatures {
			effectiveness *= GetEffectiveness(moveNature, targetNature)
		}
	}

	proficiency := 1.0
	for _, nature := range moveNatures {
		proficiency *= source.NatureDamage[nature]
		proficiency /= target.NatureResistance[nature]
	}

	return proficiency * effectiveness
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
	stat AttackStat,
	power int,
	critical float64,
	nature *NatureSet,
	random float64,
) []int {
	damages := make([]int, len(targets))
	if power == 0 {
		return damages
	}

	a_base := float64(source.Stats[BaseStat(stat)])
	a_mod := source.AttackModifiers[stat]
	attack := int(math.Floor(a_base * a_mod))

	for i, target := range targets {
		d_base := float64(target.Stats[BaseStat(stat)])
		d_mod := target.DefenseModifiers[stat]
		defense := int(math.Floor(d_base * d_mod))

		var natures []Nature
		if nature != nil {
			natures = NATURES[*nature]
		}
		nature_mod := GetNaturesModifier(source, target, natures)
		stab_mod := GetStabModifier(source, nature)

		damages[i] = DamageEquation(DamageTerms{
			Attack:   attack,
			Critical: critical,
			Defense:  defense,
			Level:    source.Level,
			Nature:   nature_mod,
			Offset:   0,
			Other:    1.00,
			Power:    power,
			Random:   random,
			STAB:     stab_mod,
		})
	}
	return damages
}
