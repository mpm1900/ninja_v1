package game

import (
	"slices"

	"github.com/google/uuid"
)

type Context struct {
	ActionID          *uuid.UUID  `json:"action_ID"`
	SourcePlayerID    *uuid.UUID  `json:"source_player_ID"`
	ParentActorID     *uuid.UUID  `json:"parent_actor_ID"`
	SourceActorID     *uuid.UUID  `json:"source_actor_ID"`
	TargetActorIDs    []uuid.UUID `json:"target_actor_IDs"`
	TargetPositionIDs []uuid.UUID `json:"target_position_IDs"`
}

func NewContext() Context {
	return Context{
		TargetActorIDs:    []uuid.UUID{},
		TargetPositionIDs: []uuid.UUID{},
	}
}

func WithTargetIDs(context Context, targetActorIDs []uuid.UUID) Context {
	c := context
	c.TargetActorIDs = targetActorIDs
	return c
}

func (g Game) GetTargets(context Context) []Actor {
	targets := g.GetActors(func(a Actor) bool {
		return slices.Contains(context.TargetActorIDs, a.ID)
	})
	posTargets := g.GetActors(func(a Actor) bool {
		if a.PositionID == nil {
			return false
		}

		return slices.Contains(context.TargetPositionIDs, *a.PositionID)
	})

	targets = append(targets, posTargets...)
	return targets
}

func (g Game) GetSource(context Context) (Actor, bool) {
	if context.SourceActorID == nil {
		return Actor{}, false
	}

	return g.GetActorByID(*context.SourceActorID)
}

func (g Game) GetParent(context Context) (Actor, bool) {
	if context.ParentActorID == nil {
		return Actor{}, false
	}

	return g.GetActorByID(*context.ParentActorID)
}
