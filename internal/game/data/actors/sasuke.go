package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Sasuke = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Sasuke Uchiha",
	Clan:         game.ClanUchiha,
	Affiliations: []string{},

	Stats: map[game.BaseStat]int{
		game.StatHP:       88,
		game.StatChakra:   75,
		game.StatNinjutsu: 120,
		game.StatGenjutsu: 120,
		game.StatTaijutsu: 90,
		game.StatSpeed:    142,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsLightning,
		game.NsYin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
