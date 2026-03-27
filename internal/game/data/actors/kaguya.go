package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Kaguya = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Kaguya Ōtsutsuki",
	Affiliations: []string{},

	Stats: map[game.BaseStat]int{
		game.StatHP:       255,
		game.StatChakra:   255,
		game.StatNinjutsu: 135,
		game.StatGenjutsu: 135,
		game.StatTaijutsu: 115,
		game.StatSpeed:    125,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYin,
		game.NsYang,
		game.NsYinYang,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
