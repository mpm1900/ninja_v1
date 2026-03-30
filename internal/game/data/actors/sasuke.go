package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Sasuke = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Sasuke Uchiha",
	SpriteURL:    "/sprites/sasuke_64.png",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffAkatsuki, game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:            80,
		game.StatStamina:       70,
		game.StatAttack:        120,
		game.StatDefense:       80,
		game.StatChakraAttack:  140,
		game.StatChakraDefense: 120,
		game.StatSpeed:         132,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsLightning,
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
