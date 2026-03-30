package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Jiraiya = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/jiraiya_64.png",
	Name:         "Jiraiya",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:            105,
		game.StatStamina:       90,
		game.StatAttack:        100,
		game.StatDefense:       113,
		game.StatChakraAttack:  100,
		game.StatChakraDefense: 119,
		game.StatSpeed:         88,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsEarth,
		game.NsYang,
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
