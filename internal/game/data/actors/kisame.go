package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Kisame = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Kisame Hoshigaki",
	SpriteURL:    "/sprites/kisame_64.png",
	Affiliations: []string{game.AffAkatsuki, game.AffKuri},

	Stats: map[game.ActorStat]int{
		game.StatHP:            120,
		game.StatStamina:       130,
		game.StatAttack:        110,
		game.StatDefense:       90,
		game.StatChakraAttack:  110,
		game.StatChakraDefense: 90,
		game.StatSpeed:         80,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWater,
	}),
	Abilities:   []game.Modifier{},
	Ability:     &modifiers.WaterAbsorb,
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Surf.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
