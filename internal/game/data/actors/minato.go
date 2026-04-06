package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Minato = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/minato_64.png",
	Name:         "Minato Namikaze",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            60,
		game.StatStamina:       80,
		game.StatAttack:        110,
		game.StatDefense:       80,
		game.StatChakraAttack:  120,
		game.StatChakraDefense: 80,
		game.StatSpeed:         200,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsWind,
		game.NsLightning,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Rasengan.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.ToadSong.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
