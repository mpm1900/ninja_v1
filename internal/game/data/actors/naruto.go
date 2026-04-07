package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Naruto = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/naruto_64.png",
	Name:         "Naruto Uzumaki",
	Clan:         game.ClanUzumaki,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            105,
		game.StatStamina:       130,
		game.StatAttack:        100,
		game.StatDefense:       80,
		game.StatChakraAttack:  105,
		game.StatChakraDefense: 105,
		game.StatSpeed:         105,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsPure,
		game.NsWind,
		game.NsYang,
	}),
	Abilities:   []game.Modifier{},
	Ability:     nil,
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Rasengan.ID,
		actions.PowerBoost.ID,
		actions.ToadSong.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
