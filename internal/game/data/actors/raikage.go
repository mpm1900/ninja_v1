package game

import (
	"ninja_v1/internal/game"
	// modifiers "ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var RaikageID = uuid.New()

func NewRaikage(playerID uuid.UUID, level int) game.Actor {

	return game.Actor{
		ID:       uuid.New(),
		PlayerID: playerID,
		ActorID:  RaikageID,
		Name:     "A (4th Raikage)",

		Level:       level,
		Experience:  0,
		ActionCount: 6,

		Active: true,
		Alive:  true,

		Stats: map[game.BaseStat]int{
			game.StatHP:       50,
			game.StatStamina:  90,
			game.StatNinjutsu: 95,
			game.StatGenjutsu: 85,
			game.StatTaijutsu: 105,
			game.StatSpeed:    180,
			game.StatEvasion:  0,
			game.StatAccuracy: 1,
		},

		Stages: map[game.BaseStat]int{
			game.StatHP:       0,
			game.StatStamina:  0,
			game.StatNinjutsu: 0,
			game.StatGenjutsu: 0,
			game.StatTaijutsu: 0,
			game.StatSpeed:    0,
			game.StatEvasion:  0,
			game.StatAccuracy: 0,
		},

		Critical: 1.50,
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
			game.NsWater,
		}),
		InnateModifiers: []game.Modifier{},
		Actions:         []game.Action{},
	}
}
