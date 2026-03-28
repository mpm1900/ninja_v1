package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Itachi = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/itachi_64.png",
	Name:         "Itachi Uchiha",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffAkatsuki, game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       65,
		game.StatChakra:   80,
		game.StatNinjutsu: 120,
		game.StatGenjutsu: 155,
		game.StatTaijutsu: 95,
		game.StatSpeed:    130,
		game.StatEvasion:  100,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsYin,
	}),

	InnateModifiers: []game.Modifier{},
	ActionCount:     6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
		actions.Recover.ID,
	},
}
