package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Kakuzu = game.ActorDef{
	ActorID:      uuid.MustParse("9a273c6c-d268-5d54-9667-4c264d5192d8"),
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
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.SearingMigraine.ID,
		actions.LeafJab.ID,
	},
}
