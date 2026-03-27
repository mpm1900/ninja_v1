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
		game.StatChakra:   60,
		game.StatNinjutsu: 95,
		game.StatGenjutsu: 80,
		game.StatTaijutsu: 105,
		game.StatSpeed:    180,
		game.StatEvasion:  0,
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
