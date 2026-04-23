package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

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
		game.StatHP:            80,
		game.StatStamina:       100,
		game.StatAttack:        76,
		game.StatDefense:       126,
		game.StatChakraAttack:  89,
		game.StatChakraDefense: 136,
		game.StatSpeed:         33,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsEarth,
		game.NsMagnet,
	}),
	Abilities: []game.Modifier{
		modifiers.SandAura,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.DragonStance.ID,
		actions.WhirlwindKick.ID,
		actions.EarthDomePrison.ID,
	},
}
