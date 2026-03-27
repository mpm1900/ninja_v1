package actors

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Orochimaru = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/orochimaru_64.png",
	Name:         "Orochimaru",
	Affiliations: []string{game.AffAkatsuki, game.AffOto},

	Stats: map[game.BaseStat]int{
		game.StatHP:       100,
		game.StatChakra:   100,
		game.StatNinjutsu: 135,
		game.StatGenjutsu: 100,
		game.StatTaijutsu: 90,
		game.StatSpeed:    101,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWind,
		game.NsEarth,
		game.NsYin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs:       []uuid.UUID{},
}
