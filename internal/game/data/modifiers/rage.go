package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var rageID = uuid.MustParse("ab1dcd38-82d6-5a88-9123-91caed67da84")
var RageTrigger game.Trigger = game.Trigger{
	ID:         uuid.MustParse("27591068-a257-5554-b654-60d8e46e30f9"),
	ModifierID: rageID,
	On:         game.OnDamageRecieve,
	Check:      game.Match__TargetActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: 0,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
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
			mutation := mutations.AddModifiers(false, AttackUpSource)
			transaction := game.MakeTransaction(mutation, mut_ctx)
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var Rage game.Modifier = game.Modifier{
	ID:          rageID,
	GroupID:     &rageID,
	Name:        "Rage",
	Description: "On damage: gain Attack Up.",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&rageID),
	},
	Triggers: []game.Trigger{
		RageTrigger,
	},
}
