package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Kaguya = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Kaguya Ōtsutsuki",
	ActionCount: 6,

	Stats: map[game.BaseStat]int{
		game.StatHP:       255,
		game.StatStamina:  255,
		game.StatNinjutsu: 125,
		game.StatGenjutsu: 125,
		game.StatTaijutsu: 115,
		game.StatSpeed:    130,
		game.StatEvasion:  0,
		game.StatAccuracy: 1,
	},
	NatureDamage: map[game.Nature]float64{
		game.NatureFire:      1.00,
		game.NatureWind:      1.00,
		game.NatureLightning: 1.00,
		game.NatureEarth:     1.00,
		game.NatureWater:     1.00,
		game.NatureYin:       1.00,
		game.NatureYang:      1.00,
	},
	NatureResistance: map[game.Nature]float64{
		game.NatureFire:      1.00,
		game.NatureWind:      1.00,
		game.NatureLightning: 1.00,
		game.NatureEarth:     1.00,
		game.NatureWater:     1.00,
		game.NatureYin:       1.00,
		game.NatureYang:      1.00,
	},

	Natures: game.MapNatures([]game.NatureSet{
		game.NsYin,
		game.NsYang,
		game.NsYinYang,
	}),
	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
