package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Itachi = game.ActorDef{
	ActorID:      uuid.MustParse("52713900-6202-50f5-9b3d-23dc408e3c63"),
	SpriteURL:    "/sprites/itachi_64.png",
	Name:         "Itachi Uchiha",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffAkatsuki, game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            60,
		game.StatStamina:       80,
		game.StatAttack:        80,
		game.StatDefense:       60,
		game.StatChakraAttack:  140,
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
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.LuckyStrikes.ID,
		actions.Disable.ID,
		actions.PatternBreak.ID,
		actions.Flash.ID,
		actions.MirageCrow.ID,
		actions.Fireball.ID,
		actions.GreatFireball.ID,
		actions.PhoenixFlower.ID,
		actions.PunishingFire.ID,
		actions.Amaterasu.ID,
		actions.Coercion.ID,
		actions.InstilFear.ID,
		actions.PerishSong.ID,
		actions.Taunt.ID,
		actions.DisarmingStrike.ID,
	},
}
