package mutations

import (
	"fmt"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func RedirectSingleTargetEnemyActions(source game.Actor) game.GameMutation {
	return game.GameMutation{
		Delta: func(g game.Game, context game.Context) game.Game {
			for i, a := range g.Actions {
				if a.Context.SourcePlayerID == nil {
					continue
				}

				targets := g.GetTargets(a.Context)
				if len(targets) != 1 {
					continue
				}

				fmt.Println("FOUND ACTIONS TO REDIRECT")

				/**
				 * One potential update here is to do a context check to make sure the new target is valid
				 */
				if *a.Context.SourcePlayerID != source.PlayerID && targets[0].PlayerID == source.PlayerID {
					if a.Mutation.TargetType == game.TargetActorID {
						g.Actions[i].Context.TargetActorIDs = []uuid.UUID{source.ID}
					}
					if a.Mutation.TargetType == game.TargetPositionID && source.PositionID != nil {
						g.Actions[i].Context.TargetPositionIDs = []uuid.UUID{*source.PositionID}
					}
				}
			}

			return g
		},
	}
}
