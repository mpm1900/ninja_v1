package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Chidori = MakeChidori()

func MakeChidori() game.Action {
	accuracy := 100
	power := 80
	recoil := 0.2
	nature := game.NsLightning
	stat := game.ChakraAttack
	config := game.ActionConfig{
		Name:     "Chidori",
		Nature:   &nature,
		Accuracy: &accuracy,
		Power:    &power,
		Stat:     &stat,
		Recoil:   &recoil,
		Jutsu:    game.Ninjutsu,
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
				damages := mutations.NewDamage(conf, game.NewDamageConfig(1, 1))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
