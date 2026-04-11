package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var sobpID = uuid.MustParse("096e2442-b231-53da-9892-91b0dea908b9")
var SealOfBodyProtectionTrigger game.Trigger = game.Trigger{
	ID:         uuid.MustParse("0a179ac1-8811-5319-9c4b-bb23c1f57ed8"),
	ModifierID: sobpID,
	On:         game.OnModifierAdd,
	Check:      game.Match__SourceActor_SourceActor, // TODO move below switches to a filter so we dont get black logs
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}

			if context.ModifierID == nil {
				return transactions
			}

			modifier, ok := g.GetModifierTxByID(*context.ModifierID)
			if !ok {
				return transactions
			}

			if modifier.Mutation.GroupID != nil {
				switch *modifier.Mutation.GroupID {
				case AttackDownID, DefenseDownID, ChakraAttackDownID, ChakraDefenseDownID, SpeedDownID, AccuracyDownID, EvasionDownID:
					mut := mutations.RemoveModifierTxByID(*context.ModifierID)
					transactions = append(transactions, game.MakeTransaction(mutations.ConsumeItem, context))
					transactions = append(transactions, game.MakeTransaction(mut, context))

				default:
					break
				}
			}

			return transactions
		},
	},
}

var SealOfBodyProtection game.Modifier = game.Modifier{
	ID:          sobpID,
	GroupID:     &sobpID,
	Name:        "Seal of Body Protection",
	Description: "On stat drop: remove it and break this seal.",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&sobpID),
	},
	Triggers: []game.Trigger{
		SealOfBodyProtectionTrigger,
	},
}
