package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var NoopAbility game.Modifier = game.Modifier{
	ID: uuid.New(),
	Name: "",
	Show: false,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{},
	Triggers: []game.Trigger{},
}

var neutralizingChakraID = uuid.MustParse("ca02f87f-a32c-4977-bc75-5720fffc0475")

var NeutralizingChakra game.Modifier = game.Modifier{
	ID:          neutralizingChakraID,
	GroupID:     &neutralizingChakraID,
	Name:        "Neutralizing Chakra",
	Description: "On enter: Disable all active abilities",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{},
	Triggers: []game.Trigger{
		{
			ID: uuid.New(),
			ModifierID: neutralizingChakraID,
			On:         game.OnActorEnter,
			Check:      game.Match__SourceActor_SourceActor,
			ActionMutation: game.ActionMutation{
				Priority: game.ActionPriorityP1,
				Filter:   game.TrueGameFilter,
				Delta: func(p, g game.Game, context game.Context) []game.GameTransaction {
					transactions := []game.GameTransaction{}
					targets := g.GetActorsFilters(context, game.ComposeAF(
						game.ActiveFilter,
						game.AliveFilter,
						game.OtherFilter,
					))
					for _, target := range targets {
						mut_ctx := context
						mut_ctx.ModifierID = &intimidateID
						mut_ctx.TargetActorIDs = []uuid.UUID{target.ID}
						mutation := game.GameMutation{
							Delta: func(p, g game.Game, context game.Context) game.Game {
								g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
									a.AuxAbility = &NoopAbility
									return a
								})
								return g
							},
						}
						transaction := game.MakeTransaction(mutation, mut_ctx)
						transactions = append(transactions, transaction)
					}

					return transactions
				},
			},
		},
	},
}
