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

	Stats: map[game.ActorStat]int{
		game.StatHP:            91,
		game.StatStamina:       90,
		game.StatAttack:        90,
		game.StatDefense:       106,
		game.StatChakraAttack:  130,
		game.StatChakraDefense: 106,
		game.StatSpeed:         77,
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
		actions.Rasengan.ID,
		actions.Haze.ID,
		actions.ToadSong.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
