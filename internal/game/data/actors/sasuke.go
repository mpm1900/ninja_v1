package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Sasuke = game.ActorDef{
	ActorID:      uuid.MustParse("2c7025a3-77c8-57b7-9120-e99b473f669f"),
	Name:         "Sasuke Uchiha",
	SpriteURL:    "/sprites/sasuke_64.png",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffAkatsuki, game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            88,
		game.StatStamina:       70,
		game.StatAttack:        100,
		game.StatDefense:       75,
		game.StatChakraAttack:  120,
		game.StatChakraDefense: 75,
		game.StatSpeed:         142,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsLightning,
	}),
	Abilities: []game.Modifier{
		modifiers.InnerFocus,
		modifiers.Unburden,
	},
	ActionCount: 4,
	ActionIDs: append([]uuid.UUID{
		actions.Chidori.ID,
		actions.ChidoriSpear.ID,
		actions.ChidoriStream.ID,
		actions.Kirin.ID,
		actions.DragonStance.ID,
		actions.DragonFire.ID,
		actions.GreatFireball.ID,
		actions.Amaterasu.ID,
	}, GlobalActions...),
}
