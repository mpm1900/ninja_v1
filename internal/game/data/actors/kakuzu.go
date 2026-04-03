package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Kakuzu = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/kakuzu_64.png",
	Name:         "Kakuzu",
	Affiliations: []string{game.AffAkatsuki, game.AffTaki},

	Stats: map[game.ActorStat]int{
		game.StatHP:            110,
		game.StatStamina:       80,
		game.StatAttack:        108,
		game.StatDefense:       125,
		game.StatChakraAttack:  120,
		game.StatChakraDefense: 79,
		game.StatSpeed:         90,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsEarth,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.SearingMigrane.ID,
		actions.LeafJab.ID,
	},
}
