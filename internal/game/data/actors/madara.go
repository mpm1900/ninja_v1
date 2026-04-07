package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Madara = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/madara_64.png",
	Name:         "Madara Uchiha",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffAkatsuki, game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            100,
		game.StatStamina:       110,
		game.StatAttack:        110,
		game.StatDefense:       80,
		game.StatChakraAttack:  165,
		game.StatChakraDefense: 105,
		game.StatSpeed:         130,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsStorm,
		game.NsYin,
	}),

	InnateModifiers: []game.Modifier{
		modifiers.AddImmunities(modifiers.AttackUpID),
	},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Glare.ID,
		actions.Protect.ID,
		actions.GreatTreeSpear.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
