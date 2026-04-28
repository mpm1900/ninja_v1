package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Raikage = game.ActorDef{
	ActorID:      uuid.MustParse("e92a1675-7ea0-52c7-a43c-f2bb2068e365"),
	SpriteURL:    "/sprites/4_raikage_64.png",
	Name:         "A (4th Raikage)",
	Affiliations: []string{game.AffKumo},

	Stats: map[game.ActorStat]int{
		game.StatHP:            80,
		game.StatStamina:       80,
		game.StatAttack:        135,
		game.StatDefense:       100,
		game.StatChakraAttack:  90,
		game.StatChakraDefense: 60,
		game.StatSpeed:         150,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsLightning,
		game.NsEarth,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: append([]uuid.UUID{
		actions.Chidori.ID,
		actions.DragonStance.ID,
		actions.Fireball.ID,
		actions.WhirlwindKick.ID,
	}, GlobalActions...),
}
