package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Kakuzu = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/kakuzu_64.png",
	Name:         "Kakuzu",
	Affiliations: []string{game.AffAkatsuki, game.AffTaki},

	Stats: map[game.BaseStat]int{
		game.StatHP:           110,
		game.StatChakra:       80,
		game.StatAttack:       108,
		game.StatDefense:      130,
		game.StatJutsu:        120,
		game.StatJutsuDefense: 70,
		game.StatSpeed:        91,
		game.StatEvasion:      100,
		game.StatAccuracy:     100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWind,
		game.NsEarth,
		game.NsWater,
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
