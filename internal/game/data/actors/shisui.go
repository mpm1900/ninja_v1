package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Shisui = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/shisui_64.png",
	Name:         "Shisui Uchiha",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            71,
		game.StatStamina:       71,
		game.StatAttack:        90,
		game.StatDefense:       90,
		game.StatChakraAttack:  155,
		game.StatChakraDefense: 112,
		game.StatSpeed:         124,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsWind,
		game.NsYin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.FollowMe.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
		actions.Recover.ID,
	},
}
