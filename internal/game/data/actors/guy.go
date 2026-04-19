package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Guy = game.ActorDef{
	ActorID:   uuid.MustParse("bf98da03-0afb-54c9-9b0a-0636552cb32d"),
	SpriteURL: "/sprites/guy_64.png",
	Name:      "Might Guy",
	Affiliations: []string{
		game.AffKonoha,
	},

	Stats: map[game.ActorStat]int{
		game.StatHP:            87,
		game.StatStamina:       80,
		game.StatAttack:        145,
		game.StatDefense:       92,
		game.StatChakraAttack:  75,
		game.StatChakraDefense: 86,
		game.StatSpeed:         115,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsTai,
		game.NsFire,
		game.NsLightning,
	}),
	Abilities: []game.Modifier{
		modifiers.Rage,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.LuckyStrikes.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
