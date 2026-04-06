package mutations

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func RedirectSingleTargetEnemyActions(source game.Actor) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			for i, a := range g.Actions {
				if a.Context.SourcePlayerID == nil {
					continue
				}

				targets := g.GetTargets(a.Context)
				if len(targets) != 1 {
					continue
				}

				if *a.Context.SourcePlayerID != source.PlayerID {
					if a.Mutation.TargetType == game.TargetActorID {
						g.Actions[i].Context.TargetActorIDs = []uuid.UUID{source.ID}
					}
					if a.Mutation.TargetType == game.TargetPositionID && source.IsActive() {
						g.Actions[i].Context.TargetPositionIDs = []uuid.UUID{*source.PositionID}
					}
				}
			}

			return g
		},
	}
}

func BoostActionPower(ratio float64) game.GameMutation {
	return game.GameMutation{
		Delta: func(p game.Game, g game.Game, context game.Context) game.Game {
			for i, a := range g.Actions {
				if a.Context.SourceActorID == nil {
					continue
				}

				targets := g.GetTargets(context)
				for _, t := range targets {

					if *a.Context.SourceActorID == t.ID && a.Mutation.Config.Power != nil {
						power := game.Round(float64(*g.Actions[i].Mutation.Config.Power) * ratio)
						g.Actions[i].Mutation.Config.Power = &power
					}
				}
			}

			return g
		},
	}
}
