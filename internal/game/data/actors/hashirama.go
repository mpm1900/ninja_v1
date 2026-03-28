package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Hashirama = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/hashirama_64.png",
	Name:         "Hashirama Senju",
	Clan:         game.ClanSenju,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       120,
		game.StatChakra:   120,
		game.StatNinjutsu: 120,
		game.StatGenjutsu: 120,
		game.StatTaijutsu: 120,
		game.StatSpeed:    120,
		game.StatEvasion:  0,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWind,
		game.NsYang,
		game.NsWood,
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
