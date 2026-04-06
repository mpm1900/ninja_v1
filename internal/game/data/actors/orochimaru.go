package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Orochimaru = game.ActorDef{
	ActorID:      uuid.New(),
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

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
