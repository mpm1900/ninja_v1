package modifiers

import (
	"ninja_v1/internal/game"
	mutations "ninja_v1/internal/game/data/game_mutations"

	"github.com/google/uuid"
)

var RageTrigger game.Trigger = game.Trigger{
	ID:    uuid.New(),
	On:    game.OnDamageRecieve,
	Check: game.MatchTargetActorIDTrigger,
	ActionMutation: game.ActionMutation{
		Priority: 0,
		Filter:   game.AllGameFilter,
		Delta: func(g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			transactions := []game.GameTransaction{}

			count, targets := g.GetTargets(context)
			if count == 0 {
				return transactions
			}

			target := targets[0]
			mut_ctx := game.Context{
				SourcePlayerID: &target.PlayerID,
				SourceActorID:  &target.ID,
				ParentActorID:  &target.ID,
			}
			mutation := mutations.AddModifiers(TaijutsuUpSource)
			transaction := game.MakeTransaction(mutation, mut_ctx)
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var Rage game.Modifier = game.Modifier{
	ID:      uuid.New(),
	GroupID: uuid.New(),
	Name:    "Rage",

	Mutations: []game.ModifierMutation{},
	Triggers: []game.Trigger{
		RageTrigger,
	},
}
