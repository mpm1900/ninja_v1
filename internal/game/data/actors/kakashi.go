package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Kakashi = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Kakashi Hatake",
	Clan:         game.ClanHatake,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatChakra:   70,
		game.StatNinjutsu: 135,
		game.StatGenjutsu: 100,
		game.StatTaijutsu: 115,
		game.StatSpeed:    120,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsLightning,
		game.NsEarth,
		game.NsYin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
