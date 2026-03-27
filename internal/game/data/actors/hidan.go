package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Hidan = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Hidan",
	Affiliations: []string{game.AffAkatsuki},

	Stats: map[game.BaseStat]int{
		game.StatHP:       190,
		game.StatChakra:   50,
		game.StatNinjutsu: 60,
		game.StatGenjutsu: 60,
		game.StatTaijutsu: 70,
		game.StatSpeed:    60,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsJashin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
