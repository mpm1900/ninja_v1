package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var consumeChakraID = uuid.New()

var ConsumeChakraTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: consumeChakraID,
	On:         game.OnKill,
	Check:      game.Match__SourceActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.SourceIsAlive,
		Delta: func(p, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			transactions := []game.GameTransaction{}

			mutation := mutations.AddModifiers(false, false, ChakraAttackUpSource)
			transaction := game.MakeTransaction(mutation, context)
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var ConsumeChakra game.Modifier = game.Modifier{
	ID:       consumeChakraID,
	GroupID:  &consumeChakraID,
	Name:     "Consume Chakra",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&consumeChakraID),
	},
	Triggers: []game.Trigger{
		ConsumeChakraTrigger,
	},
}
