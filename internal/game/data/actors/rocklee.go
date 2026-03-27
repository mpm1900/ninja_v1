package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var RockLee = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Rock Lee",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       65,
		game.StatChakra:   80,
		game.StatNinjutsu: 50,
		game.StatGenjutsu: 50,
		game.StatTaijutsu: 165,
		game.StatSpeed:    90,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsTai,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
