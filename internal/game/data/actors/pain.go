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
		game.StatHP:           80,
		game.StatChakra:       120,
		game.StatAttack:       100,
		game.StatDefense:      110,
		game.StatJutsu:        130,
		game.StatJutsuDefense: 110,
		game.StatSpeed:        110,
		game.StatEvasion:      100,
		game.StatAccuracy:     100,
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
