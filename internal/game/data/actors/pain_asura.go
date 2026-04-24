package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var PainAsura = game.ActorDef{
	ActorID:      uuid.MustParse("333051db-f9be-4d81-8520-e3553d509142"),
	SpriteURL:    "/sprites/pain_asura_64.png",
	Name:         "Pain (Asura Path)",
	Affiliations: []string{game.AffAkatsuki, game.AffAme},
	Stats: map[game.ActorStat]int{
		game.StatHP:            65,
		game.StatStamina:       100,
		game.StatAttack:        125,
		game.StatDefense:       50,
		game.StatChakraAttack:  125,
		game.StatChakraDefense: 60,
		game.StatSpeed:         115,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYinYang,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Sekiryoku.ID,
		actions.BodyReplacement.ID,
		actions.BlackNeedle.ID,
		actions.RetreatingStrike.ID,
	},
	Immunities: map[uuid.UUID]struct{}{
		modifiers.FeelingOfPain.ID: {},
		modifiers.VoiceOfPain.ID:   {},
		modifiers.VesselOfPain.ID:  {},
	},
}
