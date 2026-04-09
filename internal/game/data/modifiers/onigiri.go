package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var onigiriID = uuid.MustParse("0f737c71-12e9-5dae-94ce-9fbe02a2a2db")
var OnigiriIDTrigger game.Trigger = game.Trigger{
	ID:         uuid.MustParse("59a4d8f3-18df-5300-a3ff-d5c2c61bbfec"),
	ModifierID: onigiriID,
	On:         game.OnDamageRecieve,
	Check:      game.ComposeTF(game.Match__TargetActor_SourceActor, game.Source__IsAtOrBelowHealth(0.5)),
	ActionMutation: game.ActionMutation{
		Priority: 0,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			transactions := []game.GameTransaction{}

			targets := g.GetTargets(context)
			for _, target := range targets {
				mut_ctx := game.MakeContextForActor(target)

				heal := mutations.RatioHeal(0.25)
				heal_tx := game.MakeTransaction(heal, mut_ctx)
				consume_tx := game.MakeTransaction(mutations.ConsumeItem, mut_ctx)
				transactions = append(transactions, consume_tx)
				transactions = append(transactions, heal_tx)
			}

			return transactions
		},
	},
}

var Onigiri game.Modifier = game.Modifier{
	ID:       onigiriID,
	GroupID:  &onigiriID,
	Name:     "Onigiri",
	Show:     true,
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&onigiriID),
	},
	Triggers: []game.Trigger{
		OnigiriIDTrigger,
	},
}
