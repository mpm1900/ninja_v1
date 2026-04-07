package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Choji = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/choji_64.png",
	Name:         "Choji Akimichi",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            150,
		game.StatStamina:       100,
		game.StatAttack:        100,
		game.StatDefense:       115,
		game.StatChakraAttack:  65,
		game.StatChakraDefense: 65,
		game.StatSpeed:         35,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsEarth,
		game.NsYang,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
