package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Tsunade = game.ActorDef{
	ActorID:      uuid.MustParse("aad15064-9fcb-4bdb-a66b-36e56694316e"),
	SpriteURL:    "/sprites/tsunade_64.png",
	Name:         "Tsunade",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            110,
		game.StatStamina:       100,
		game.StatAttack:        100,
		game.StatDefense:       110,
		game.StatChakraAttack:  75,
		game.StatChakraDefense: 120,
		game.StatSpeed:         75,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWater,
		game.NsYang,
	}),
	Abilities: []game.Modifier{
		modifiers.Regeneration,
	},
	ActionCount: 4,
	ActionIDs: append([]uuid.UUID{
		actions.Haze.ID,
		actions.Fireball.ID,
		actions.HeavyPunch.ID,
		actions.SageMode.ID,
		actions.TeamHeal.ID,
	}, GlobalActions...),
}
