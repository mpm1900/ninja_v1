package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Sasuke = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Sasuke Uchiha",
	ActionCount: 6,

	Stats: map[game.BaseStat]int{
		game.StatHP:       88,
		game.StatStamina:  75,
		game.StatNinjutsu: 120,
		game.StatGenjutsu: 120,
		game.StatTaijutsu: 90,
		game.StatSpeed:    142,
		game.StatEvasion:  0,
		game.StatAccuracy: 1,
	},
	NatureDamage: map[game.Nature]float64{
		game.NatureFire:      1.30,
		game.NatureWind:      1.00,
		game.NatureLightning: 1.05,
		game.NatureEarth:     0.95,
		game.NatureWater:     0.90,
		game.NatureYin:       1.00,
		game.NatureYang:      1.00,
	},
	NatureResistance: map[game.Nature]float64{
		game.NatureFire:      1.15,
		game.NatureWind:      1.00,
		game.NatureLightning: 1.00,
		game.NatureEarth:     1.00,
		game.NatureWater:     0.95,
		game.NatureYin:       1.00,
		game.NatureYang:      1.00,
	},
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsLightning,
		game.NsYin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
