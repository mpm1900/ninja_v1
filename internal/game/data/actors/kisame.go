package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Kisame = game.ActorDef{
	ActorID:      uuid.New(),
	Name:         "Kisame Hoshigaki",
	SpriteURL:    "/sprites/kisame_64.png",
	Affiliations: []string{game.AffAkatsuki, game.AffKuri},

	Stats: map[game.BaseStat]int{
		game.StatHP:       110,
		game.StatChakra:   130,
		game.StatNinjutsu: 110,
		game.StatGenjutsu: 60,
		game.StatTaijutsu: 110,
		game.StatSpeed:    100,
		game.StatEvasion:  100,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsPure,
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
