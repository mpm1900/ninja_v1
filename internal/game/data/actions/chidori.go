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
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherTeamFilter, game.ActiveFilter, game.AliveFilter),
		ContextValidate: game.PositionsLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.AllGameFilter,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				damages := mutations.NewDamage(config, game.NewDamageConfig())
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
