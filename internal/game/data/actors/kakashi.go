package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var Kakashi = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/kakashi_64.png",
	Name:         "Kakashi Hatake",
	Clan:         game.ClanHatake,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.BaseStat]int{
		game.StatHP:       85,
		game.StatChakra:   60,
		game.StatNinjutsu: 135,
		game.StatGenjutsu: 105,
		game.StatTaijutsu: 120,
		game.StatSpeed:    120,
		game.StatEvasion:  100,
		game.StatAccuracy: 100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsLightning,
		game.NsEarth,
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
	},
}
