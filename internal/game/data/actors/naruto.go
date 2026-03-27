package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Naruto = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/naruto_64.png",
	Name:         "Naruto Uzumaki (Toad Sage)",
	Clan:         game.ClanUzumaki,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       105,
		game.StatChakra:   130,
		game.StatNinjutsu: 105,
		game.StatGenjutsu: 75,
		game.StatTaijutsu: 100,
		game.StatSpeed:    105,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsPure,
		game.NsWind,
		game.NsYang,
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
