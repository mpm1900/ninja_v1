package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var regnerationID = uuid.MustParse("23d1b13b-4ad6-464e-a005-936cbc121ae1")
var RegnerationTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: regnerationID,
	On:         game.OnActorLeave,
	Check:      game.Match__SourceActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			source, ok := g.GetSource(context)
			if !ok {
				return transactions
			}

			mut := game.RatioHeal(0.25)
			mut_ctx := game.MakeContextForActor(source)
			mut_tx := game.MakeTransaction(mut, mut_ctx)
			transactions = append(transactions, mut_tx)

			return transactions
		},
	},
}

var Regneration game.Modifier = game.Modifier{
	ID:          regnerationID,
	GroupID:     &regnerationID,
	Name:        "Regneration",
	Description: "On exit: heal for 1/4th HP.",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&regnerationID),
	},
	Triggers: []game.Trigger{
		RegnerationTrigger,
	},
}
