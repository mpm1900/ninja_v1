package game

import (
	"math"
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
	level_mod := float64(2*terms.Level)/5 + 2
	pow_ad := float64(terms.Power) * float64(terms.Attack) / float64(terms.Defense)
	base := (level_mod*pow_ad)/50 + 2
	raw := (base * terms.Critical * terms.Nature * terms.Random * terms.Other) + float64(terms.Offset)
	return int(math.Floor(raw))
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

func GetDamage(source ResolvedActor, targets []ResolvedActor, stat AttackStat, power int, natures []Nature) []int {
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

		nature_mod := GetNaturesModifier(source, target, natures)

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
			STAB:     1.00,
		})
	}
	return damages
}
