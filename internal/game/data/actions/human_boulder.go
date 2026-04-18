package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var HumanBoulder = MakeHumanBoulder()

func MakeHumanBoulder() game.Action {
	config := game.ActionConfig{
		Name:        "Human Boulder",
		Description: "Damage is based of the user's Defense stat rather than Attack.",
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(80),
		Stat:        game.Ptr(game.StatDefense),
		Nature:      game.Ptr(game.NsTai),
		Jutsu:       game.Taijutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}
	return game.Action{
		ID:              uuid.MustParse("05b5376a-5c76-4f72-bc2c-c148ad068e40"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, 1))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
