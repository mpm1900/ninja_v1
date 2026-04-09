package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var intimidateID = uuid.MustParse("0a6d78f6-c09b-5463-ad56-e9d549fc7ca9")
var IntimidateTrigger game.Trigger = game.Trigger{
	ID:         uuid.MustParse("4af16c70-9c55-5f4c-a5b5-6a24397167c5"),
	ModifierID: intimidateID,
	On:         game.OnActorEnter,
	Check:      game.Match__SourceActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			targets := g.GetActorsFilters(context, game.ComposeAF(
				game.ActiveFilter,
				game.AliveFilter,
				game.OtherFilter,
				// game.OtherTeamFilter,
			))

			for _, target := range targets {
				mut_ctx := game.MakeContextForActor(target)
				mutation := mutations.AddModifiers(false, false, AttackDownSource)
				transaction := game.MakeTransaction(mutation, mut_ctx)
				transactions = append(transactions, transaction)
			}

			return transactions
		},
	},
}

var Intimidate game.Modifier = game.Modifier{
	ID:       intimidateID,
	GroupID:  &intimidateID,
	Name:     "Intimidate",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&intimidateID),
	},
	Triggers: []game.Trigger{
		IntimidateTrigger,
	},
}
