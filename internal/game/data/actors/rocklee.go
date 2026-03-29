package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var RockLee = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/rocklee_64.png",
	Name:         "Rock Lee",
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       70,
		game.StatChakra:   80,
		game.StatNinjutsu: 50,
		game.StatGenjutsu: 50,
		game.StatTaijutsu: 145,
		game.StatSpeed:    105,
		game.StatEvasion:  100,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsTai,
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
