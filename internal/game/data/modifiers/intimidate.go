package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var intimidateID = uuid.New()
var IntimidateTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: intimidateID,
	On:         game.OnActorEnter,
	Check:      game.Match__SourceActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: 0,
		Filter:   game.TrueGameFilter,
		Delta: func(g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}
			targets := g.GetActorsFilters(context, game.ComposeAF(
				game.ActiveFilter,
				game.AliveFilter,
				game.OtherTeamFilter,
			))

			for _, target := range targets {
				mut_ctx := game.Context{
					SourcePlayerID: &target.PlayerID,
					SourceActorID:  &target.ID,
					ParentActorID:  &target.ID,
				}
				mutation := mutations.AddModifiers(AttackDownSource)
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
	Duration: game.ModifierDurationInf,
	Mutations: []game.ModifierMutation{
		game.NewNoopSource(&intimidateID),
	},
	Triggers: []game.Trigger{
		IntimidateTrigger,
	},
}
