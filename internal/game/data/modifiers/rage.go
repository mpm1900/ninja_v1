package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var rageID = uuid.New()
var RageTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: rageID,
	On:         game.OnDamageRecieve,
	Check:      game.Match__TargetActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: 0,
		Filter:   game.AllGameFilter,
		Delta: func(g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			transactions := []game.GameTransaction{}

			targets := g.GetTargets(context)
			if len(targets) == 0 {
				return transactions
			}

			target := targets[0]
			mut_ctx := game.Context{
				SourcePlayerID: &target.PlayerID,
				SourceActorID:  &target.ID,
				ParentActorID:  &target.ID,
			}
			mutation := mutations.AddModifiers(AttackUpSource)
			transaction := game.MakeTransaction(mutation, mut_ctx)
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var RageGroupId = uuid.New()
var Rage game.Modifier = game.Modifier{
	ID:       rageID,
	GroupID:  RageGroupId,
	Name:     "Rage",
	Duration: game.ModifierDurationInf,
	Mutations: []game.ActorMutation{
		game.NewNoop(&RageGroupId),
	},
	Triggers: []game.Trigger{
		RageTrigger,
	},
}
