package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var PainNaraka = game.ActorDef{
	ActorID:      uuid.MustParse("e8906381-8f07-4034-a235-1e7153f2da54"),
	SpriteURL:    "/sprites/pain_naraka_64.png",
	Name:         "Pain (Naraka Path)",
	Affiliations: []string{game.AffAkatsuki, game.AffAme},
	Stats: map[game.ActorStat]int{
		game.StatHP:            65,
		game.StatStamina:       100,
		game.StatAttack:        110,
		game.StatDefense:       90,
		game.StatChakraAttack:  80,
		game.StatChakraDefense: 90,
		game.StatSpeed:         45,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYinYang,
	}),
	Abilities: []game.Modifier{
		modifiers.NeutralizingChakra,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.PerishSong.ID,
		actions.TempleOfNirvana.ID,
		actions.Flash.ID,
		actions.Revival.ID,
		actions.RetreatingStrike.ID,
		actions.SummonAlly.ID,
		actions.BlackNeedle.ID,
	},
	Immunities: map[uuid.UUID]struct{}{
		modifiers.FeelingOfPain.ID: {},
		modifiers.VoiceOfPain.ID:   {},
		modifiers.VesselOfPain.ID:  {},
	},
}
