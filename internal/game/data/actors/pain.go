package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Pain = game.ActorDef{
	ActorID:      uuid.MustParse("f344914c-26d5-5e7b-8673-ad73d5b8f334"),
	SpriteURL:    "/sprites/pain_64.png",
	Name:         "Pain",
	Clan:         game.ClanUzumaki,
	Affiliations: []string{game.AffAkatsuki, game.AffAme},

	Stats: map[game.ActorStat]int{
		game.StatHP:            106,
		game.StatStamina:       120,
		game.StatAttack:        90,
		game.StatDefense:       100,
		game.StatChakraAttack:  130,
		game.StatChakraDefense: 144,
		game.StatSpeed:         100,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYinYang,
	}),
	Abilities: []game.Modifier{
		modifiers.VesselOfPain,
		modifiers.SpeedBoost,
		modifiers.Raincaller,
	},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		actions.Sekiryoku.ID,
		actions.Protect.ID,
		actions.Tailwind.ID,
		actions.LeafJab.ID,
		actions.Surf.ID,
	},
}
