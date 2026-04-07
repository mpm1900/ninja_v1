package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Onoki = game.ActorDef{
	ActorID:   uuid.New(),
	SpriteURL: "/sprites/onoki_64.png",
	Name:      "Ōnoki",
	Affiliations: []string{
		game.AffIwa,
	},

	Stats: map[game.ActorStat]int{
		game.StatHP:            56,
		game.StatStamina:       80,
		game.StatAttack:        80,
		game.StatDefense:       114,
		game.StatChakraAttack:  124,
		game.StatChakraDefense: 60,
		game.StatSpeed:         136,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsEarth,
		game.NsWind,
		game.NsParticle,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.LuckyStrikes.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
