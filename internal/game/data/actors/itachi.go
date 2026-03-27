package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Itachi = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Itachi Uchiha",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffAkatsuki, game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       70,
		game.StatChakra:   80,
		game.StatNinjutsu: 120,
		game.StatGenjutsu: 155,
		game.StatTaijutsu: 90,
		game.StatSpeed:    135,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsYin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
