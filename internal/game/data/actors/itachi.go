package game

import (
	"ninja_v1/internal/game"
	// modifiers "ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var ItachiID = uuid.New()

func NewItachi(playerID uuid.UUID, level int) game.Actor {

	return game.Actor{
		ID:       uuid.New(),
		PlayerID: playerID,
		ActorID:  ItachiID,
		Name:     "Itachi Uchiha",

		Level:       level,
		Experience:  0,
		ActionCount: 6,

		Active: true,
		Alive:  true,

		Stats: map[game.BaseStat]int{
			game.StatHP:       60,
			game.StatStamina:  60,
			game.StatNinjutsu: 120,
			game.StatGenjutsu: 155,
			game.StatTaijutsu: 90,
			game.StatSpeed:    125,
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
			game.NatureFire:      1.30,
			game.NatureWind:      1.00,
			game.NatureLightning: 1.05,
			game.NatureEarth:     0.95,
			game.NatureWater:     0.90,
			game.NatureYin:       1.00,
			game.NatureYang:      1.00,
		},
		NatureResistance: map[game.Nature]float64{
			game.NatureFire:      1.15,
			game.NatureWind:      1.00,
			game.NatureLightning: 1.00,
			game.NatureEarth:     1.00,
			game.NatureWater:     0.95,
			game.NatureYin:       1.00,
			game.NatureYang:      1.00,
		},

		Natures: game.MapNatures([]game.NatureSet{
			game.NsFire,
			game.NsYin,
		}),

		InnateModifiers: []game.Modifier{},
		Actions:         []game.Action{},
	}
}
