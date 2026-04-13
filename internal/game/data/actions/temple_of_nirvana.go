package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var TempleOfNirvana = MakeTempleOfNirvana()

func MakeTempleOfNirvana() game.Action {
	ID := uuid.MustParse("d59535f2-9cb5-4268-854e-4d9a1d6b7c70")
	nature := game.NsYin

	config := game.ActionConfig{
		Name:        "Temple Of Nirvana",
		Description: "Puts target to sleep",
		Nature:      &nature,
		TargetCount: game.Ptr(1),
		Accuracy:    game.Ptr(100),
		Cost:        game.Ptr(30),
		Jutsu:       game.Genjutsu,
	}

	return game.Action{
		ID:              ID,
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(*config.Cost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter: game.ComposeGF(
				game.SourceIsAlive,
			),
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				targets := g.GetTargets(context)
				for _, target := range targets {

					mut_ctx := game.Context{
						SourcePlayerID: &target.PlayerID,
						SourceActorID:  &target.ID,
						ParentActorID:  nil, // do not remove on switch
						TargetActorIDs: []uuid.UUID{target.ID},
					}
					mutation := mutations.Sleep
					transaction := game.MakeTransaction(mutation, mut_ctx)
					transactions = append(transactions, transaction)
				}

				return transactions
			},
		},
	}
}
