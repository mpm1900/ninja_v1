package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Kakuzu = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Kakuzu",
	ActionCount: 6,
	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatStamina:  110,
		game.StatNinjutsu: 134,
		game.StatGenjutsu: 85,
		game.StatTaijutsu: 110,
		game.StatSpeed:    86,
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
		game.NsFire,
		game.NsWind,
		game.NsEarth,
	}),

	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
