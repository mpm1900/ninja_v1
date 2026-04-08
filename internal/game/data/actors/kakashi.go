package actors

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"
	"ninja_v1/internal/game/data/modifiers"

	"github.com/google/uuid"
)

func getKakashiResistances() map[game.Nature]float64 {
	res := game.NewNatureSetValues()
	res[game.NatureLightning] = 2.0
	res[game.NatureEarth] = 0.5
	return res
}

var Kakashi = game.ActorDef{
	ActorID:      uuid.New(),
	SpriteURL:    "/sprites/kakashi_64.png",
	Name:         "Kakashi Hatake",
	Clan:         game.ClanHatake,
	Affiliations: []string{game.AffKonoha},
	Stats: map[game.ActorStat]int{
		game.StatHP:            95,
		game.StatStamina:       70,
		game.StatAttack:        111,
		game.StatDefense:       80,
		game.StatChakraAttack:  135,
		game.StatChakraDefense: 80,
		game.StatSpeed:         109,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: getKakashiResistances(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsLightning,
		game.NsEarth,
		game.NsYin,
	}),
	Abilities: []game.Modifier{
		modifiers.AccuracyUpSource,
	},
	ActionCount: 6,
	ActionIDs: []uuid.UUID{
		game.Switch.ID,
		actions.Chidori.ID,
		actions.DragonDance.ID,
		actions.Fireball.ID,
		actions.LeafJab.ID,
	},
}
