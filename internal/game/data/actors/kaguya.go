package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Kaguya = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Kaguya Ōtsutsuki",
	Affiliations: []string{},

	Stats: map[game.BaseStat]int{
		game.StatHP:           255,
		game.StatChakra:       255,
		game.StatAttack:       135,
		game.StatDefense:      255,
		game.StatJutsu:        135,
		game.StatJutsuDefense: 255,
		game.StatSpeed:        125,
		game.StatEvasion:      100,
		game.StatAccuracy:     100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsYin,
		game.NsYang,
		game.NsYinYang,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
