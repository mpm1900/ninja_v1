package game

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Jiraiya = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Jiraiya",
	ActionCount: 6,

	Stats: map[game.BaseStat]int{
		game.StatHP:       95,
		game.StatStamina:  108,
		game.StatNinjutsu: 140,
		game.StatGenjutsu: 109,
		game.StatTaijutsu: 100,
		game.StatSpeed:    88,
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
		game.NsEarth,
		game.NsYang,
	}),

	InnateModifiers: []game.Modifier{},
	ActionIDs:       []uuid.UUID{},
}
