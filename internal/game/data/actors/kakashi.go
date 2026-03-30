package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Kakashi = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/kakashi_64.png",
	Name:         "Kakashi Hatake",
	Clan:         game.ClanHatake,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:           85,
		game.StatChakra:       70,
		game.StatAttack:       125,
		game.StatDefense:      80,
		game.StatJutsu:        140,
		game.StatJutsuDefense: 110,
		game.StatSpeed:        120,
		game.StatEvasion:      100,
		game.StatAccuracy:     100,
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
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
