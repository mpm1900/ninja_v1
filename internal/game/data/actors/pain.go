package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Pain = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/pain_64.png",
	Name:         "Pain",
	Clan:         game.ClanUzumaki,
	Affiliations: []string{game.AffAkatsuki, game.AffAme},

	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatChakra:   120,
		game.StatNinjutsu: 144,
		game.StatGenjutsu: 130,
		game.StatTaijutsu: 100,
		game.StatSpeed:    106,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYin,
		game.NsYang,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
