package game

import (
	"ninja_v1/internal/game"
	// modifiers "ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var YamatoID = uuid.New()

func NewYamato(playerID uuid.UUID, level int) game.Actor {

	return game.Actor{
		ID:       uuid.New(),
		PlayerID: playerID,
		ActorID:  YamatoID,
		Name:     "Yamato",

		Level:       level,
		Experience:  0,
		ActionCount: 6,

		Active: true,
		Alive:  true,

		Stats: map[game.BaseStat]int{
			game.StatHP:       76,
			game.StatStamina:  71,
			game.StatNinjutsu: 104,
			game.StatGenjutsu: 71,
			game.StatTaijutsu: 104,
			game.StatSpeed:    108,
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
			game.NsEarth,
			game.NsYang,
			game.NsWood,
		}),
		InnateModifiers: []game.Modifier{},
		Actions:         []game.Action{},
	}
}
