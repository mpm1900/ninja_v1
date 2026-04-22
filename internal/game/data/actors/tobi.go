package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Tobi = game.ActorDef{
	ActorID:      uuid.MustParse("e01398df-0a65-45a4-a4d6-878af3fa9d4c"),
	SpriteURL:    "/sprites/tobi_64.png",
	Name:         "Tobi",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffAkatsuki, game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            100,
		game.StatStamina:       100,
		game.StatAttack:        135,
		game.StatDefense:       120,
		game.StatChakraAttack:  60,
		game.StatChakraDefense: 85,
		game.StatSpeed:         50,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsYin,
	}),
	Abilities:   []game.Modifier{},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.BodyReplacement.ID,
		actions.KamuiCounter.ID,
		actions.KamuiSlash.ID,
		actions.SageMode.ID,
		actions.PatternBreak.ID,
		actions.KusariChains.ID,
		actions.SwordsStance.ID,
	},
}
