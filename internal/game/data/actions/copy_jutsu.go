package actions

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var CopyJutsu = MakeCopyJutsu()

func MakeCopyJutsu() game.Action {
	config := game.ActionConfig{
		Name:        "Copy Jutsu",
		Nature:      game.Ptr(game.NsYin),
		Jutsu:       game.Ninjutsu,
		Description: "Copies target's last used action.",
	}
	return game.Action{
		ID:              uuid.MustParse("7bc916c9-f762-472c-ac29-d5e7996b64e3"),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.ComposeAF(game.TeamFilter, game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				source, ok := g.GetSource(context)
				if !ok {
					return transactions
				}

				targets := g.GetTargets(context)
				for _, target := range targets {
					if target.LastUsedActionID == nil {
						continue
					}

					resolved := target.Resolve(g)
					action, ok := resolved.GetActionByID(*resolved.LastUsedActionID)
					if !ok {
						continue
					}

					action_ctx := game.Context{
						ActionID:       &action.ID,
						SourceActorID:  &source.ID,
						SourcePlayerID: &source.PlayerID,
						TargetActorIDs: []uuid.UUID{target.ID}, // this assignment could be problematic in the future. maybe.
					}

					transactions = append(transactions, action.Delta(p, g, action_ctx)...)
					// TODO, maybe add the action to use until switch out? (modifier add action)
					//transaction := game.MakeTransaction(mut, context)
					//transactions = append(transactions, transaction)
				}

				return transactions
			},
		},
	}
}
