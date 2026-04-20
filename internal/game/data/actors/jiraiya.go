package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Jiraiya = game.ActorDef{
	ActorID:      uuid.MustParse("fbf88375-8f7c-5380-97b9-2f2748581e76"),
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
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Rasengan.ID,
		actions.GiantRasengan.ID,
		actions.Haze.ID,
		actions.ToadSong.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
