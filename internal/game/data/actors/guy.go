package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Guy = game.ActorDef{
	ActorID:     uuid.New(),
	Name:        "Might Guy",
	ActionCount: 6,

	Stats: map[game.BaseStat]int{
		game.StatHP:       70,
		game.StatStamina:  120,
		game.StatNinjutsu: 100,
		game.StatGenjutsu: 55,
		game.StatTaijutsu: 140,
		game.StatSpeed:    115,
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
		game.NsTai,
		game.NsFire,
		game.NsLightning,
	}),
	InnateModifiers: []game.Modifier{
		modifiers.Rage,
	},
	ActionIDs: []uuid.UUID{
		actions.LeafJab.ID,
	},
}
