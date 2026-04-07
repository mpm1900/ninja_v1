package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Tobirama = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/tobirama_64.png",
	Name:         "Tobirama Senju",
	Clan:         game.ClanSenju,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            70,
		game.StatStamina:       100,
		game.StatAttack:        110,
		game.StatDefense:       75,
		game.StatChakraAttack:  145,
		game.StatChakraDefense: 85,
		game.StatSpeed:         145,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWater,
		game.NsLightning,
		game.NsYin,
	}),
	Abilities:   []game.Modifier{},
	Ability:     &modifiers.ConsumeChakra,
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Protect.ID,
		actions.Tailwind.ID,
		actions.LeechSeed.ID,
		actions.Haze.ID,
		actions.GreatTreeSpear.ID,
		actions.LeafJab.ID,
		actions.Substitution.ID,
	},
}
