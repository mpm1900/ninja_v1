package actions

import (
	"fmt"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

var Substitution = MakeSubstitution()

func MakeSubstitution() game.Action {
	nature := game.NsYin
	config := game.ActionConfig{
		Name:        "Substitution",
		Nature:      &nature,
		Jutsu:       game.Ninjutsu,
		Description: "Summons a substitute to take damage. Pay 1/4th of Max HP.",
	}

	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.NoneFilter,
		ContextValidate: game.TargetLengthFilter(0),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.ComposeGF(game.SourceIsAlive, game.SourceHasHpRatio(0.25)),
			Delta: func(p, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
				transactions := []game.GameTransaction{}
				s, ok := g.GetSource(context)
				if !ok {
					return transactions
				}
				source := s.Resolve(g)

				mut := game.GameMutation{
					Delta: func(mp, mg game.Game, mc game.Context) game.Game {
						mg.UpdateActor(*mc.SourceActorID, func(a game.Actor) game.Actor {
							hp := game.Round(float64(source.Stats[game.StatHP]) * 0.25)
							summon := game.MakeActor(game.ActorDef{
								ActorID: uuid.New(),
								SpriteURL: "/sprites/sub_64.png",
								Name:    fmt.Sprintf("Substitute (%d)", hp),
								Stats: map[game.ActorStat]int{
									game.StatHP: hp,
								},
							}, a.PlayerID, a.Experience, []uuid.UUID{}, map[uuid.UUID]game.Action{})
							a.SetSummonFromActor(&summon, true)
							return a
						})
						return mg
					},
				}

				transactions = append(transactions, game.MakeTransaction(mut, context))

				return transactions
			},
		},
	}
}
