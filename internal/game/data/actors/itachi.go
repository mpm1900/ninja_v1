package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Itachi = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/itachi_64.png",
	Name:         "Itachi Uchiha",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffAkatsuki, game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            60,
		game.StatStamina:       80,
		game.StatAttack:        80,
		game.StatDefense:       60,
		game.StatChakraAttack:  145,
		game.StatChakraDefense: 130,
		game.StatSpeed:         130,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsYin,
	}),
	Abilities: []game.Modifier{
		modifiers.Intimidate,
	},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.FollowMe.ID,
		actions.Glare.ID,
		actions.MirageCrow.ID,
		actions.Fireball.ID,
		actions.Amaterasu.ID,
		actions.Recover.ID,
	},
}
