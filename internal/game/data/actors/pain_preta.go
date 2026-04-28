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
		game.StatHP:            250,
		game.StatStamina:       100,
		game.StatAttack:        10,
		game.StatDefense:       10,
		game.StatChakraAttack:  80,
		game.StatChakraDefense: 105,
		game.StatSpeed:         55,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYinYang,
	}),
	Abilities: []game.Modifier{
		modifiers.ChainsOfPain,
		modifiers.NeutralizingChakra,
	},
	ActionCount: 4,
	ActionIDs: append([]uuid.UUID{
		actions.Recover.ID,
		actions.RetreatingStrike.ID,
		actions.NegateJutsu.ID,
		actions.CopyJutsu.ID,
		actions.ChannelChakra.ID,
		actions.Expansion.ID,
		actions.BlackNeedle.ID,
	}, GlobalActions...),
	Immunities: map[uuid.UUID]struct{}{
		modifiers.BurdenOfPain.ID:    {},
		modifiers.ChainsOfPain.ID:    {},
		modifiers.JudgementOfPain.ID: {},
		modifiers.VoiceOfPain.ID:     {},
	},
}
