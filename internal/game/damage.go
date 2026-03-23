package game

import "math"

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

func DamageEquation(config DamageTerms) int {
	level_mod := float64(2*config.Level)/5 + 2
	pow_ad := float64(config.Attack) / float64(config.Defense)
	term := (level_mod*pow_ad)/50 + 2
	raw := term*config.Critical*config.Nature*config.Random*config.Other + float64(config.Offset)
	return int(math.Floor(raw))
}

func GetDamage(source ResolvedActor, targets []ResolvedActor, stat BaseStat, power int) []int {
	damages := make([]int, len(targets))
	for i, target := range targets {
		damages[i] = DamageEquation(DamageTerms{
			Attack:   source.Stats[stat],
			Critical: 1.00,
			Defense:  target.Stats[stat],
			Level:    source.Level,
			Nature:   1.00,
			Offset:   0,
			Other:    1.00,
			Power:    power,
			Random:   1.00,
			STAB:     1.00,
		})
	}
	return damages
}
