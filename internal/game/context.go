package game

import (
	"slices"

	"github.com/google/uuid"
)

type Context struct {
	SourcePlayerID    *uuid.UUID  `json:"source_player_ID"`
	ParentActorID     *uuid.UUID  `json:"parent_actor_ID"`
	SourceActorID     *uuid.UUID  `json:"source_actor_ID"`
	TargetActorIDs    []uuid.UUID `json:"target_actor_IDs"`
	TargetPositionIDs []uuid.UUID `json:"target_position_IDs"`
}

func GetTargets(g Game, context Context) []Actor {
	count := len(context.TargetActorIDs) + len(context.TargetPositionIDs)
	targets := make([]Actor, 0, count)
	for _, targetID := range context.TargetActorIDs {
		i := slices.IndexFunc(g.Actors, func(a Actor) bool { return a.ID == targetID })
		if i == -1 {
			continue
		}

		targets = append(targets, g.Actors[i])
	}
	for _, positionID := range context.TargetPositionIDs {
		i := slices.IndexFunc(g.Actors, func(a Actor) bool {
			return a.PositionID != nil && *a.PositionID == positionID
		})
		if i == -1 {
			break
		}

		targets = append(targets, g.Actors[i])
	}

	return targets
}
