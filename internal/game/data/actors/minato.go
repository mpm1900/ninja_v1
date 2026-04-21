package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Minato = game.ActorDef{
	ActorID:      uuid.MustParse("dce0fef4-265e-5dc8-9d81-700ed3fc4877"),
	SpriteURL:    "/sprites/minato_64.png",
	Name:         "Minato Namikaze",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            80,
		game.StatStamina:       80,
		game.StatAttack:        130,
		game.StatDefense:       80,
		game.StatChakraAttack:  130,
		game.StatChakraDefense: 80,
		game.StatSpeed:         100,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWind,
		game.NsLightning,
	}),
	Abilities: []game.Modifier{
		modifiers.SpeedBoost,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Rasengan.ID,
		actions.GiantRasengan.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.SageMode.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
		actions.FlyingRaijin.ID,
		actions.SummonAlly.ID,
	},
}
