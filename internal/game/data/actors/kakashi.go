package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Kakashi = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Kakashi Hatake",
	ActionCount: 6,

	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatChakra:   70,
		game.StatNinjutsu: 130,
		game.StatGenjutsu: 100,
		game.StatTaijutsu: 130,
		game.StatSpeed:    120,
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
		game.NsLightning,
		game.NsEarth,
		game.NsYin,
	}),
	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
