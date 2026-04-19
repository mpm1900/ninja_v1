package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Hidan = game.ActorDef{
	ActorID:      uuid.MustParse("735c88e7-9c5a-5a99-9605-0feeea1ccdb5"),
	SpriteURL:    "/sprites/hidan_64.png",
	Name:         "Hidan",
	Affiliations: []string{game.AffAkatsuki, game.AffYuga},

	Stats: map[game.ActorStat]int{
		game.StatHP:            200,
		game.StatStamina:       100,
		game.StatAttack:        70,
		game.StatDefense:       55,
		game.StatChakraAttack:  30,
		game.StatChakraDefense: 55,
		game.StatSpeed:         70,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsJashin,
	}),
	Abilities: []game.Modifier{
		modifiers.ShadowCage,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Distraction.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.BloodPrice.ID,
		actions.LeafJab.ID,
		actions.Curse.ID,
		actions.PerishSong.ID,
	},
}
