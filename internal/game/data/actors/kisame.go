package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Kisame = game.ActorDef{
	ActorID:      uuid.MustParse("e34e2dec-6b2b-59f5-92c4-afb7e473f3e9"),
	Name:         "Kisame Hoshigaki",
	SpriteURL:    "/sprites/kisame_64.png",
	Affiliations: []string{game.AffAkatsuki, game.AffKuri},

	Stats: map[game.ActorStat]int{
		game.StatHP:            120,
		game.StatStamina:       130,
		game.StatAttack:        110,
		game.StatDefense:       90,
		game.StatChakraAttack:  110,
		game.StatChakraDefense: 90,
		game.StatSpeed:         80,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWater,
	}),
	Abilities: []game.Modifier{
		modifiers.WaterAbsorb,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Surf.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.WhirlwindKick.ID,
		actions.HiddenMist.ID,
	},
}
