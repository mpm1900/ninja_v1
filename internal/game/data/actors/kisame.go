package game

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Kisame = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Kisame Hoshigaki",
	ActionCount: 6,

	Stats: map[game.BaseStat]int{
		game.StatHP:       110,
		game.StatStamina:  130,
		game.StatNinjutsu: 110,
		game.StatGenjutsu: 60,
		game.StatTaijutsu: 110,
		game.StatSpeed:    95,
		game.StatEvasion:  0,
		game.StatAccuracy: 1,
	},
	NatureDamage: map[game.Nature]float64{
		game.NatureFire:      0.90,
		game.NatureWind:      0.95,
		game.NatureLightning: 0.85,
		game.NatureEarth:     1.00,
		game.NatureWater:     1.35,
		game.NatureYin:       1.00,
		game.NatureYang:      1.00,
	},
	NatureResistance: map[game.Nature]float64{
		game.NatureFire:      1.15,
		game.NatureWind:      1.00,
		game.NatureLightning: 0.85,
		game.NatureEarth:     1.00,
		game.NatureWater:     1.30,
		game.NatureYin:       1.00,
		game.NatureYang:      1.00,
	},
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWater,
		game.NsYang,
	}),

	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
