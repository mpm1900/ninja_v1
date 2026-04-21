package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var PainHuman = game.ActorDef{
	ActorID:      uuid.MustParse("81803b0a-fb3a-4a79-bd8a-b4531503bbed"),
	SpriteURL:    "/sprites/pain_human_64.png",
	Name:         "Pain (Human Path)",
	Affiliations: []string{game.AffAkatsuki, game.AffAme},
	Stats: map[game.ActorStat]int{
		game.StatHP:            65,
		game.StatStamina:       75,
		game.StatAttack:        75,
		game.StatDefense:       130,
		game.StatChakraAttack:  75,
		game.StatChakraDefense: 130,
		game.StatSpeed:         95,
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
		actions.PerishSong.ID,
		actions.MindTransfer.ID,
		actions.TempleOfNirvana.ID,
		actions.Flash.ID,
	},
}
