package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Yamato = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Yamato",
	ActionCount: 6,

	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatStamina:  101,
		game.StatNinjutsu: 120,
		game.StatGenjutsu: 70,
		game.StatTaijutsu: 91,
		game.StatSpeed:    84,
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
		game.NsEarth,
		game.NsWater,
		game.NsWood,
	}),
	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
