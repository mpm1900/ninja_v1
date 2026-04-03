package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Pain = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/pain_64.png",
	Name:         "Pain",
	Clan:         game.ClanUzumaki,
	Affiliations: []string{game.AffAkatsuki, game.AffAme},

	Stats: map[game.ActorStat]int{
		game.StatHP:            80,
		game.StatStamina:       120,
		game.StatAttack:        100,
		game.StatDefense:       110,
		game.StatChakraAttack:  130,
		game.StatChakraDefense: 110,
		game.StatSpeed:         110,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYin,
		game.NsYang,
	}),

	InnateModifiers: []game.Modifier{
		modifiers.VesselOfPain,
	},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Sekiryoku.ID,
		actions.Protect.ID,
		actions.Tailwind.ID,
		actions.LeafJab.ID,
	},
}
