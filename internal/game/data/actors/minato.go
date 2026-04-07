package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Minato = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/minato_64.png",
	Name:         "Minato Namikaze",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            80,
		game.StatStamina:       80,
		game.StatAttack:        130,
		game.StatDefense:       80,
		game.StatChakraAttack:  130,
		game.StatChakraDefense: 80,
		game.StatSpeed:         100,
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

	InnateModifiers: []game.Modifier{
		modifiers.SpeedBoost,
	},
	ActionCount: 6,
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
