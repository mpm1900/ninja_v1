package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var brittleSashID = uuid.New()
var BrittleSashTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: brittleSashID,
	On:         game.OnImmortalSave,
	Check:      game.ComposeTF(game.Match__SourceActor_SourceActor),
	ActionMutation: game.ActionMutation{
		Priority: 0,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			transactions := []game.GameTransaction{}

			source, ok := g.GetSource(context)
			if !ok {
				return transactions
			}

			mut_ctx := game.MakeContextForActor(source)
			consume_tx := game.MakeTransaction(mutations.ConsumeItem, mut_ctx)
			transactions = append(transactions, consume_tx)
			return transactions
		},
	},
}

var BrittleSash game.Modifier = game.Modifier{
	ID:       brittleSashID,
	GroupID:  &brittleSashID,
	Name:     "Brittle Sash",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.MakeActorMutation(
			&brittleSashID,
			game.MutPriorityDefault,
			game.ComposeAF(game.SourceFilter, game.ActiveFilter, game.FullHealthFilter),
			func(g game.Game, a game.Actor, ctx game.Context) game.Actor {
				a.Immortal = true
				return a
			}),
	},
	Triggers: []game.Trigger{
		BrittleSashTrigger,
	},
}
