package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Yamato = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Yamato",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatChakra:   101,
		game.StatNinjutsu: 120,
		game.StatGenjutsu: 70,
		game.StatTaijutsu: 91,
		game.StatSpeed:    84,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsEarth,
		game.NsWater,
		game.NsWood,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
