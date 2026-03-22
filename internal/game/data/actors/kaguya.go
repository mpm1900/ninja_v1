package game

import (
	"ninja_v1/internal/game"
	// modifiers "ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var KaguyaID = uuid.New()

func NewKaguya(playerID uuid.UUID, level int) game.Actor {

	return game.Actor{
		ID:       uuid.New(),
		PlayerID: playerID,
		ActorID:  KaguyaID,
		Name:     "Kaguya Ōtsutsuki",

		Level:       level,
		Experience:  0,
		ActionCount: 6,

		Active: true,
		Alive:  true,

		Stats: map[game.BaseStat]int{
			game.StatHP:       255,
			game.StatStamina:  255,
			game.StatNinjutsu: 125,
			game.StatGenjutsu: 125,
			game.StatTaijutsu: 115,
			game.StatSpeed:    130,
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
			game.NsYin,
			game.NsYang,
			game.NsYinYang,
		}),
		InnateModifiers: []game.Modifier{},
		Actions:         []game.Action{},
	}
}
