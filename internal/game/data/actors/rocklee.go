package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var RockLee = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/rocklee_64.png",
	Name:         "Rock Lee",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            70,
		game.StatStamina:       80,
		game.StatAttack:        150,
		game.StatDefense:       90,
		game.StatChakraAttack:  20,
		game.StatChakraDefense: 70,
		game.StatSpeed:         105,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsTai,
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
