package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Naruto = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Naruto Uzumaki (Toad Sage)",
	ActionCount: 6,
	Stats: map[game.BaseStat]int{
		game.StatHP:       105,
		game.StatChakra:   130,
		game.StatNinjutsu: 105,
		game.StatGenjutsu: 75,
		game.StatTaijutsu: 100,
		game.StatSpeed:    105,
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
		game.NsPure,
		game.NsWind,
		game.NsYang,
	}),

	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
