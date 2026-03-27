package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Minato = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Minato Namikaze",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       60,
		game.StatChakra:   60,
		game.StatNinjutsu: 120,
		game.StatGenjutsu: 80,
		game.StatTaijutsu: 100,
		game.StatSpeed:    200,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsWind,
		game.NsLightning,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
