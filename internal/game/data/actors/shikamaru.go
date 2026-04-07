package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Shikamaru = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/shikamaru_64.png",
	Name:         "Shikamaru Nara",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            70,
		game.StatStamina:       100,
		game.StatAttack:        55,
		game.StatDefense:       65,
		game.StatChakraAttack:  95,
		game.StatChakraDefense: 105,
		game.StatSpeed:         85,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsEarth,
		game.NsYin,
	}),
	Abilities:   []game.Modifier{},
	Ability:     &modifiers.FastThinking,
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Haze.ID,
		actions.Tailwind.ID,
		actions.FollowMe.ID,
	},
}
