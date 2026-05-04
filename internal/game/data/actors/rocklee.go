package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var RockLee = game.ActorDef{
	ActorID:      uuid.MustParse("f7189c34-fb21-54d6-be14-28a473d36c53"),
	SpriteURL:    "/sprites/rocklee_64.png",
	Name:         "Rock Lee",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            60,
		game.StatStamina:       100,
		game.StatAttack:        60,
		game.StatDefense:       75,
		game.StatChakraAttack:  60,
		game.StatChakraDefense: 75,
		game.StatSpeed:         100,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsTai,
	}),
	Abilities: []game.Modifier{
		modifiers.PurePower,
	},
	ActionCount: 4,
	ActionIDs: append([]uuid.UUID{
		actions.LuckyStrikes.ID,
		actions.DragonStance.ID,
		actions.WhirlwindKick.ID,
		actions.HeavyPunch.ID,
		actions.FlyingLotus.ID,
		actions.Asakujaku.ID,
		actions.Hirudora.ID,
	}, GlobalActions...),
}
