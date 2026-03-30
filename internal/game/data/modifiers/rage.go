package modifiers

import (
	"ninja_v1/internal/game"
	actor_mutations "ninja_v1/internal/game/data/actor_mutations"
	game_mutations "ninja_v1/internal/game/data/game_mutations"

	"github.com/google/uuid"
)

var RageTrigger game.Trigger = game.Trigger{
	ID:    uuid.New(),
	On:    game.OnDamageRecieve,
	Check: game.Match__TargetActor_SourceActor,
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
			mutation := game_mutations.AddModifiers(AttackUpSource)
			transaction := game.MakeTransaction(mutation, mut_ctx)
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var RageGroupId = uuid.New()
var Rage game.Modifier = game.Modifier{
	ID:      uuid.New(),
	GroupID: RageGroupId,
	Name:    "Rage",
	Mutations: []game.ModifierMutation{
		actor_mutations.NewNoop(&RageGroupId),
	},
	Triggers: []game.Trigger{
		RageTrigger,
	},
}
