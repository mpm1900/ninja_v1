package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var PainPreta = game.ActorDef{
	ActorID:      uuid.MustParse("276059f9-61e2-440c-8e0b-2b2b29e88268"),
	SpriteURL:    "/sprites/pain_preta_64.png",
	Name:         "Pain (Preta Path)",
	Affiliations: []string{game.AffAkatsuki, game.AffAme},
	Stats: map[game.ActorStat]int{
		game.StatHP:            98,
		game.StatStamina:       100,
		game.StatAttack:        53,
		game.StatDefense:       85,
		game.StatChakraAttack:  87,
		game.StatChakraDefense: 105,
		game.StatSpeed:         67,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYinYang,
	}),
	Abilities: []game.Modifier{
		modifiers.StatusReflection,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Recover.ID,
		actions.Tailwind.ID,
		actions.PerishSong.ID,
		actions.Expansion.ID,
		actions.Taunt.ID,
		actions.RetreatingStrike.ID,
		actions.BodyReplacement.ID,
		actions.CopyJutsu.ID,
		actions.Fireball.ID,

		actions.TempleOfNirvana.ID,
	},
}
