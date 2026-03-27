package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Jiraiya = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Jiraiya",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       95,
		game.StatChakra:   90,
		game.StatNinjutsu: 140,
		game.StatGenjutsu: 109,
		game.StatTaijutsu: 100,
		game.StatSpeed:    88,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsEarth,
		game.NsYang,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
