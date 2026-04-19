package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Deidara = game.ActorDef{
	ActorID:      uuid.MustParse("d069a1f9-376e-56ce-9edf-0924a6fed8f1"),
	SpriteURL:    "/sprites/deidara_64.png",
	Name:         "Deidara",
	Affiliations: []string{game.AffAkatsuki, game.AffIwa},

	Stats: map[game.ActorStat]int{
		game.StatHP:            100,
		game.StatStamina:       110,
		game.StatAttack:        60,
		game.StatDefense:       80,
		game.StatChakraAttack:  130,
		game.StatChakraDefense: 110,
		game.StatSpeed:         95,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsLightning,
		game.NsEarth,
		game.NsExplosion,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
		actions.C0UltimateArt.ID,
		actions.C1Bird.ID,
	},
}
