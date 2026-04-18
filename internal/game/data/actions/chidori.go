package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Chidori = MakeChidori()

func MakeChidori() game.Action {
	config := game.ActionConfig{
		Name:       "Chidori",
		Nature:     game.Ptr(game.NsLightning),
		Accuracy:   game.Ptr(100),
		Power:      game.Ptr(80),
		Stat:       game.Ptr(game.StatChakraAttack),
		Recoil:     game.Ptr(0.2),
		Jutsu:      game.Ninjutsu,
		CritChance: game.Ptr(5),
		CritMod:    1.5,
	}
	return game.Action{
		ID:              uuid.MustParse("c1502330-764c-56f8-9c9e-f41b933a90f0"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherTeamFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor()))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
