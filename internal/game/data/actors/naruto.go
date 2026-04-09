package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Naruto = game.ActorDef{
	ActorID:      uuid.MustParse("7b8d8818-ebb3-5c79-8d67-20c5df3d026d"),
	SpriteURL:    "/sprites/naruto_64.png",
	Name:         "Naruto Uzumaki",
	Clan:         game.ClanUzumaki,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            105,
		game.StatStamina:       130,
		game.StatAttack:        100,
		game.StatDefense:       80,
		game.StatChakraAttack:  105,
		game.StatChakraDefense: 105,
		game.StatSpeed:         105,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsPure,
		game.NsWind,
		game.NsYang,
	}),
	Abilities: []game.Modifier{
		modifiers.FirstAid,
	},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		actions.Rasengan.ID,
		actions.PowerBoost.ID,
		actions.ToadSong.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
