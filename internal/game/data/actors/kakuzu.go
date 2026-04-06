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
		game.StatHP:            80,
		game.StatStamina:       80,
		game.StatAttack:        110,
		game.StatDefense:       125,
		game.StatChakraAttack:  125,
		game.StatChakraDefense: 80,
		game.StatSpeed:         80,
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
