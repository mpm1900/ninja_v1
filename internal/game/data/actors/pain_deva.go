package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var PainDeva = game.ActorDef{
	ActorID:      uuid.MustParse("58e03e29-77bd-494f-a3d3-6a555519eced"),
	SpriteURL:    "/sprites/pain_deva_64.png",
	Name:         "Pain (Deva Path)",
	Affiliations: []string{game.AffAkatsuki, game.AffAme},
	Stats: map[game.ActorStat]int{
		game.StatHP:            105,
		game.StatStamina:       100,
		game.StatAttack:        115,
		game.StatDefense:       85,
		game.StatChakraAttack:  90,
		game.StatChakraDefense: 75,
		game.StatSpeed:         100,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYinYang,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Sekiryoku.ID,
		actions.BodyReplacement.ID,
		actions.Tailwind.ID,
		actions.MindTransfer.ID,
	},
}
