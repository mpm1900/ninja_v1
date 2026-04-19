package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

var Shisui = game.ActorDef{
	ActorID:      uuid.MustParse("3b9a402f-9ac1-5e47-8d56-37ec1ed287ff"),
	SpriteURL:    "/sprites/shisui_64.png",
	Name:         "Shisui Uchiha",
	Clan:         game.ClanUchiha,
	Affiliations: []string{game.AffKonoha},

	Stats: map[game.ActorStat]int{
		game.StatHP:            100,
		game.StatStamina:       70,
		game.StatAttack:        90,
		game.StatDefense:       75,
		game.StatChakraAttack:  133,
		game.StatChakraDefense: 75,
		game.StatSpeed:         127,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsFire,
		game.NsWind,
		game.NsYin,
	}),
	Abilities: []game.Modifier{
		modifiers.PriorityFailure,
	},
	ActionCount: 4,
	ActionIDs: []uuid.UUID{
		actions.Distraction.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.BodyFlicker.ID,
		actions.Recover.ID,
	},
}
