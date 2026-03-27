package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Deidara = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/deidara_64.png",
	Name:         "Deidara",
	Affiliations: []string{game.AffAkatsuki, game.AffIwa},

	Stats: map[game.BaseStat]int{
		game.StatHP:       85,
		game.StatChakra:   110,
		game.StatNinjutsu: 140,
		game.StatGenjutsu: 80,
		game.StatTaijutsu: 80,
		game.StatSpeed:    110,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsLightning,
		game.NsEarth,
		game.NsExplosion,
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
