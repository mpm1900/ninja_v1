package game

import (
	"ninja_v1/internal/game"
	// modifiers "ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var KisameID = uuid.New()

func NewKisame(playerID uuid.UUID, level int) game.Actor {

	return game.Actor{
		ID:       uuid.New(),
		PlayerID: playerID,
		ActorID:  KisameID,
		Name:     "Kisame Hoshigaki",

		Level:       level,
		Experience:  0,
		ActionCount: 6,

		Active: true,
		Alive:  true,

		Stats: map[game.BaseStat]int{
			game.StatHP:       110,
			game.StatStamina:  120,
			game.StatNinjutsu: 110,
			game.StatGenjutsu: 60,
			game.StatTaijutsu: 110,
			game.StatSpeed:    90,
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
			game.NatureFire:      0.90,
			game.NatureWind:      0.95,
			game.NatureLightning: 0.85,
			game.NatureEarth:     1.00,
			game.NatureWater:     1.35,
			game.NatureYin:       1.00,
			game.NatureYang:      1.00,
		},
		NatureResistance: map[game.Nature]float64{
			game.NatureFire:      1.15,
			game.NatureWind:      1.00,
			game.NatureLightning: 0.85,
			game.NatureEarth:     1.00,
			game.NatureWater:     1.30,
			game.NatureYin:       1.00,
			game.NatureYang:      1.00,
		},

		Natures: []game.NatureSet{
			game.NsWater,
			game.NsYang,
		},

		InnateModifiers: []game.Modifier{},
		Actions:         []game.Action{},
	}
}
