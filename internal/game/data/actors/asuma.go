package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Asuma = game.ActorDef{
	ActorID:      uuid.MustParse("592da549-ce61-453c-9857-27f012a65ad9"),
	SpriteURL:    "/sprites/asuma_64.png",
	Name:         "Asuma Sarutobi",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            90,
		game.StatStamina:       90,
		game.StatAttack:        110,
		game.StatDefense:       80,
		game.StatChakraAttack:  110,
		game.StatChakraDefense: 80,
		game.StatSpeed:         95,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsWind,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Haze.ID,
		actions.Fireball.ID,
		actions.DragonFire.ID,
		actions.Firestorm.ID,
		actions.WhirlwindKick.ID,
		actions.HeavyPunch.ID,
	},
}
