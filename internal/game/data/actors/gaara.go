package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Gaara = game.ActorDef{
	ActorID:   uuid.MustParse("b32fadc0-a6c4-4cf3-8f34-f91f67eb1b77"),
	SpriteURL: "/sprites/gaara_64.png",
	Name:      "Gaara",
	Affiliations: []string{
		game.AffSun,
	},

	Stats: map[game.ActorStat]int{
		game.StatHP:            70,
		game.StatStamina:       100,
		game.StatAttack:        85,
		game.StatDefense:       145,
		game.StatChakraAttack:  60,
		game.StatChakraDefense: 55,
		game.StatSpeed:         65,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWind,
		game.NsEarth,
		game.NsLightning,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		actions.LuckyStrikes.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
