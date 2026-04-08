package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var bpsID = uuid.New()
var BodyProtectionSealTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: bpsID,
	On:         game.OnModifierAdd,
	Check:      game.Match__SourceActor_SourceActor,
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

var BodyProtectionSeal game.Modifier = game.Modifier{
	ID:       bpsID,
	GroupID:  &bpsID,
	Name:     "Body Protection Seal",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&bpsID),
	},
	Triggers: []game.Trigger{
		BodyProtectionSealTrigger,
	},
}
