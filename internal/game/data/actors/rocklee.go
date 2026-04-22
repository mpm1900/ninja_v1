package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var RockLee = game.ActorDef{
	ActorID:      uuid.MustParse("f7189c34-fb21-54d6-be14-28a473d36c53"),
	SpriteURL:    "/sprites/rocklee_64.png",
	Name:         "Rock Lee",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            50,
		game.StatStamina:       80,
		game.StatAttack:        150,
		game.StatDefense:       80,
		game.StatChakraAttack:  20,
		game.StatChakraDefense: 70,
		game.StatSpeed:         105,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsTai,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.LuckyStrikes.ID,
		actions.DragonDance.ID,
		actions.WhirlwindKick.ID,
		actions.HeavyPunch.ID,
		actions.Asakujaku.ID,
	},
}
