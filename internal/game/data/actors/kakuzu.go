package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Kakuzu = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Kakuzu",
	Affiliations: []string{game.AffAkatsuki},

	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatChakra:   110,
		game.StatNinjutsu: 134,
		game.StatGenjutsu: 85,
		game.StatTaijutsu: 110,
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
