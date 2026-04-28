package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Raikage4 = game.ActorDef{
	ActorID:      uuid.MustParse("aecc5cde-93ae-42f7-8e0f-0e9db4294c6a"),
	SpriteURL:    "/sprites/4_raikage_64.png",
	Name:         "Ay (4th Raikage)",
	Affiliations: []string{game.AffKumo},

	Stats: map[game.ActorStat]int{
		game.StatHP:            70,
		game.StatStamina:       100,
		game.StatAttack:        130,
		game.StatDefense:       100,
		game.StatChakraAttack:  70,
		game.StatChakraDefense: 60,
		game.StatSpeed:         150,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsLightning,
		game.NsWater,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: append([]uuid.UUID{
		actions.DragonStance.ID,
		actions.RockFist.ID,
		actions.LightningLariat.ID,
	}, GlobalActions...),
}
