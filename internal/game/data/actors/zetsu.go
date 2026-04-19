package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Zetsu = game.ActorDef{
	ActorID:      uuid.MustParse("dfdd143c-faba-4ab2-b607-cbe69baf68e7"),
	SpriteURL:    "/sprites/zetsu_64.png",
	Name:         "Zetsu",
	Affiliations: []string{game.AffAkatsuki},

	Stats: map[game.ActorStat]int{
		game.StatHP:            114,
		game.StatStamina:       100,
		game.StatAttack:        85,
		game.StatDefense:       70,
		game.StatChakraAttack:  85,
		game.StatChakraDefense: 85,
		game.StatSpeed:         30,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYin,
		game.NsYang,
		game.NsWood,
	}),
	Abilities: []game.Modifier{
		modifiers.Regneration,
	},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		actions.BodyReplacement.ID,
		actions.TempleOfNirvana.ID,
		actions.Distraction.ID,
		actions.Graft.ID,
	},
}
