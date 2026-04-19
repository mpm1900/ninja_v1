package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Orochimaru = game.ActorDef{
	ActorID:      uuid.MustParse("2e9d220b-be84-524b-b7a3-b078c226fa2d"),
	SpriteURL:    "/sprites/orochimaru_64.png",
	Name:         "Orochimaru",
	Affiliations: []string{game.AffAkatsuki, game.AffOto},

	Stats: map[game.ActorStat]int{
		game.StatHP:            92,
		game.StatStamina:       90,
		game.StatAttack:        105,
		game.StatDefense:       90,
		game.StatChakraAttack:  125,
		game.StatChakraDefense: 90,
		game.StatSpeed:         98,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWind,
		game.NsEarth,
		game.NsYin,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
		actions.MindTransfer.ID,
	},
}
