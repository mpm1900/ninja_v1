package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Hidan = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Hidan",
	ActionCount: 6,

	Stats: map[game.BaseStat]int{
		game.StatHP:       190,
		game.StatChakra:   60,
		game.StatNinjutsu: 58,
		game.StatGenjutsu: 58,
		game.StatTaijutsu: 80,
		game.StatSpeed:    80,
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
		game.NsJashin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
