package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Raikage = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/4_raikage_64.png",
	Name:         "A (4th Raikage)",
	Affiliations: []string{game.AffKumo},

	Stats: map[game.BaseStat]int{
		game.StatHP:       90,
		game.StatChakra:   80,
		game.StatNinjutsu: 105,
		game.StatGenjutsu: 70,
		game.StatTaijutsu: 120,
		game.StatSpeed:    180,
		game.StatEvasion:  100,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsLightning,
		game.NsEarth,
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
