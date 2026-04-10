package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Gamabunta game.ActorDef = game.ActorDef{
	ActorID:   uuid.New(),
	SpriteURL: "/sprites/sub_64.png",
	Name:      "Gamabunta",
	Stats: map[game.ActorStat]int{
		game.StatHP:            106,
		game.StatStamina:       100,
		game.StatAttack:        90,
		game.StatDefense:       130,
		game.StatChakraAttack:  90,
		game.StatChakraDefense: 154,
		game.StatSpeed:         110,
		game.StatEvasion:       100,
		game.StatAccuracy:      100,
	},
	NatureDamage:     game.NewNatureSetValues(),
	NatureResistance: game.NewNatureSetValues(),
	Natures: game.MapNatures([]game.NatureSet{
		game.NsWater,
	}),
}

var GamabuntaActions []game.Action = []game.Action{
	Surf,
}

var SummonGamabunta = MakeSummonGamabunta()

func MakeSummonGamabunta() game.Action {
	nature := game.NsYang
	config := game.ActionConfig{
		Name:        "Summon Gamabunta",
		Nature:      &nature,
		Jutsu:       game.Ninjutsu,
		Description: "",
		Cooldown:    game.Ptr(10),
	}

	return game.Action{
		ID:              uuid.MustParse("17967cc9-5b82-43d4-9cd6-3a4a2c70ca39"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.ComposeGF(game.SourceIsAlive, game.SourceHasHpRatio(0.25)),
			Delta: func(p, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
				transactions := []game.GameTransaction{}

				mut := game.GameMutation{
					Delta: func(mp, mg game.Game, mc game.Context) game.Game {
						mg.UpdateActor(*mc.SourceActorID, func(a game.Actor) game.Actor {
							summon := game.MakeActor(
								Gamabunta,
								a.PlayerID,
								a.Experience,
								nil,
								nil,
								GamabuntaActions,
								map[game.ActorStat]int{},
							)
							a.SetSummonFromActor(&summon, false)
							return a
						})
						return mg
					},
				}

				transactions = append(
					transactions,
					game.MakeTransaction(mut, context),
				)

				return transactions
			},
		},
	}
}
