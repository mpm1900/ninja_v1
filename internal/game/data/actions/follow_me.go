package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var FollowMe = MakeFollowMe()

func MakeFollowMe() game.Action {
	nature := game.NsYin
	config := game.ActionConfig{
		Name:        "Follow Me",
		Nature:      &nature,
		Jutsu:       game.Genjutsu,
		Description: "Changes the target of enemy actions to this user if able.",
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityP2,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
				game.SourceIsActionOnCooldown,
			),
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				source, ok := g.GetSource(context)
				if !ok {
					return transactions
				}

				mutation := mutations.RedirectSingleTargetEnemyActions(source)
				transaction := game.MakeTransaction(mutation, context)
				transactions = append(transactions, transaction)

				return transactions
			},
		},
	}
}
