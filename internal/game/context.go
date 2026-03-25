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

func (g Game) GetTargets(context Context) (int, []Actor) {
	targets := g.GetActors(func(a Actor) bool {
		return slices.Contains(context.TargetActorIDs, a.ID)
	})
	posTargets := g.GetActors(func(a Actor) bool {
		if a.State.PositionID == nil {
			return false
		}

		return slices.Contains(context.TargetPositionIDs, *a.State.PositionID)
	})

	targets = append(targets, posTargets...)
	return len(targets), targets
}

func (g Game) GetSource(context Context) (bool, Actor) {
	if context.SourceActorID == nil {
		return false, Actor{}
	}

	return g.GetActor(func(a Actor) bool {
		return a.ID == *context.SourceActorID
	})
}

func (g Game) GetParent(context Context) (bool, Actor) {
	if context.ParentActorID == nil {
		return false, Actor{}
	}

	return g.GetActor(func(a Actor) bool {
		return a.ID == *context.ParentActorID
	})
}
