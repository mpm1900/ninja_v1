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

func DamageEquation(terms DamageTerms) int {
	pow_ad := float64(terms.Power) * float64(terms.Attack) / float64(terms.Defense)
	level_mod := float64(2*terms.Level)/5 + 2
	base := (pow_ad*level_mod)/50 + 2
	raw := (base * terms.Critical * terms.Nature * terms.Random * terms.Other)
	fmt.Printf(
		"(((%d * %d / %d) * ((2 * %d / 5) + 2) / 50 + 2) * (%f) = %f \n",
		terms.Power, terms.Attack, terms.Defense, terms.Level, terms.Nature, raw,
	)
	return int(math.Floor(raw)) + terms.Offset
}

func GetNaturesModifier(source, target ResolvedActor, natures []Nature) float64 {
	nature_mod := 1.0
	for _, nature := range natures {
		a_offset := (source.NatureDamage[nature] - 1)
		d_offset := (target.NatureResistance[nature] - 1)
		nature_mod += (a_offset - d_offset)
	}

	return nature_mod
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

func GetDamage(source ResolvedActor, targets []ResolvedActor, stat AttackStat, power int, nature *NatureSet) []int {
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
			Critical: 1.00,
			Defense:  defense,
			Level:    source.Level,
			Nature:   nature_mod,
			Offset:   0,
			Other:    1.00,
			Power:    power,
			Random:   1.00,
			STAB:     stab_mod,
		})
	}
	return damages
}
