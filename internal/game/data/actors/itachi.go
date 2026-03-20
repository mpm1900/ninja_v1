package game

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var ItachiID = uuid.New()

func NewItachi(playerID uuid.UUID, level int) game.Actor {
	modifier := game.MakeModifier("Ninjusu doubler")
	modifier.Mutations = []game.ModifierMutation{
		{
			ModifierID: &modifier.ID,
			ActorMutation: game.ActorMutation{
				Filter: func(actor game.Actor, context *game.Context) bool {
					return actor.ID == context.SourceActorID
				},
				Delta: func(actor game.Actor, context *game.Context) game.Actor {
					actor.Stats[game.StatNinjutsu] = actor.Stats[game.StatNinjutsu] * 2
					return actor
				},
			},
		},
	}
	return game.Actor{
		ID:       uuid.New(),
		PlayerID: playerID,
		ActorID:  ItachiID,
		Name:     "Itachi",

		Level:       level,
		Experience:  0,
		ActionCount: 6,

		Active: true,
		Alive:  true,

		Stats: map[game.BaseStat]int{
			game.StatHP:       55,
			game.StatStamina:  80,
			game.StatNinjutsu: 120,
			game.StatGenjutsu: 155,
			game.StatTaijutsu: 80,
			game.StatSpeed:    120,
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
			game.NatureYin:       1.35,
			game.NatureYang:      0.95,
		},
		NatureResistance: map[game.Nature]float64{
			game.NatureFire:      1.15,
			game.NatureWind:      1.00,
			game.NatureLightning: 1.00,
			game.NatureEarth:     1.00,
			game.NatureWater:     0.95,
			game.NatureYin:       1.25,
			game.NatureYang:      0.90,
		},

		Natures: []game.NatureSet{
			game.NsFire,
			game.NsYin,
		},

		InnateModifiers: []game.Modifier{modifier},
		Actions:         []game.Action[game.Game]{},
	}
}
