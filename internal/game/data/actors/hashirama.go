package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Hashirama = game.ActorDef{
	ActorID:      uuid.MustParse("6955b478-2b59-520a-afa3-e995d3cba9e9"),
	SpriteURL:    "/sprites/hashirama_64.png",
	Name:         "Hashirama Senju",
	Clan:         game.ClanSenju,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            100,
		game.StatStamina:       120,
		game.StatAttack:        155,
		game.StatDefense:       150,
		game.StatChakraAttack:  105,
		game.StatChakraDefense: 130,
		game.StatSpeed:         50,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWind,
		game.NsWood,
	}),
	Abilities: []game.Modifier{
		modifiers.Regeneration,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.BodyReplacement.ID,
		actions.Tailwind.ID,
		actions.LeechSeed.ID,
		actions.Haze.ID,
		actions.GreatTreeSpear.ID,
		actions.WhirlwindKick.ID,
		actions.ShadowClone.ID,
		actions.HeavyPunch.ID,
		actions.SageMode.ID,
	},
}
