package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Kakuzu = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Kakuzu",
	Affiliations: []string{game.AffAkatsuki, game.AffTaki},

	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatChakra:   110,
		game.StatNinjutsu: 128,
		game.StatGenjutsu: 85,
		game.StatTaijutsu: 116,
		game.StatSpeed:    86,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsWind,
		game.NsEarth,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
