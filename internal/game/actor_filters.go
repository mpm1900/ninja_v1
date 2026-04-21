package game

import (
	"slices"

	"github.com/google/uuid"
)

/**
 * Actor Filters
 */

type ActorFilter func(Game, Actor, Context) bool

func ComposeAF(filters ...ActorFilter) ActorFilter {
	return func(game Game, actor Actor, context Context) bool {
		for _, filter := range filters {
			if !filter(game, actor, context) {
				return false
			}
		}

		return true
	}
}
func AllFilter(game Game, actor Actor, context Context) bool {
	return true
}
func NoneFilter(game Game, actor Actor, context Context) bool {
	return false
}
func OtherFilter(game Game, actor Actor, context Context) bool {
	if context.SourceActorID == nil {
		return false
	}
	return actor.ID != *context.SourceActorID
}

/**
 * This filter doesn't need to be a resolved filter
 * actor.Alive is a special case were modifiers don't modify it, we just mutate it
 * so this is a safe check to make without resoloving the Actor to a ResolvedActor
 */
func AliveFilter(game Game, actor Actor, context Context) bool {
	return actor.Alive
}
func NotAliveFilter(game Game, actor Actor, context Context) bool {
	return !actor.Alive
}
func FullHealthFilter(game Game, actor Actor, context Context) bool {
	return actor.Damage == 0
}
func ActiveFilter(game Game, actor Actor, context Context) bool {
	return actor.IsActive()
}
func InactiveFilter(game Game, actor Actor, context Context) bool {
	return actor.PositionID == nil
}
func SourceFilter(game Game, actor Actor, context Context) bool {
	if context.SourceActorID == nil {
		return false
	}
	return actor.ID == *context.SourceActorID
}
func ParentFilter(game Game, actor Actor, context Context) bool {
	if context.ParentActorID == nil {
		return false
	}
	return actor.ID == *context.ParentActorID
}
func TargetableFilter(game Game, actor Actor, context Context) bool {
	return ComposeAF(
		AliveFilter,
		ActiveFilter,
		func(game Game, actor Actor, context Context) bool {
			resolved := actor.Resolve(game)
			if resolved.State == ActorStateIncorporeal {
				return false
			}
			return true
		},
	)(game, actor, context)
}
func TargetFilter(game Game, actor Actor, context Context) bool {
	if slices.Contains(context.TargetActorIDs, actor.ID) {
		return true
	}
	if actor.IsActive() && slices.Contains(context.TargetPositionIDs, *actor.PositionID) {
		return true
	}

	return false
}
func TeamFilter(game Game, actor Actor, context Context) bool {
	if context.SourcePlayerID == nil {
		return false
	}
	return actor.PlayerID == *context.SourcePlayerID
}
func OtherTeamFilter(game Game, actor Actor, context Context) bool {
	if context.SourcePlayerID == nil {
		return false
	}
	return actor.PlayerID != *context.SourcePlayerID
}
func IsAtOrBelowHealthRatio(ratio float64) func(Game, Actor, Context) bool {
	return func(game Game, actor Actor, context Context) bool {
		hp := float64(actor.Stats[StatHP])
		damage := float64(actor.Damage)
		return ratio >= (hp-damage)/hp
	}
}
func HasAppliedModifier(modifierID uuid.UUID) func(Game, Actor, Context) bool {
	return func(game Game, actor Actor, context Context) bool {
		for mid, _ := range actor.AppliedModifiers {
			if mid == modifierID {
				return true
			}
		}

		return false
	}
}

/**
 * RESOLVED FILTERS
 * These filters required modifiers to be resolved to check things like
 * protected, chakra amount, and things that can change with modifers
 * THESE CANNOT BE USED IN MODIFIERS
 */
func RHasChakraFilter(amount int) func(Game, Actor, Context) bool {
	return func(game Game, actor Actor, context Context) bool {
		resolved := actor.Resolve(game)
		return resolved.HasChakra(amount)
	}
}
func RIsProtectedFilter(protected bool) func(Game, Actor, Context) bool {
	return func(game Game, actor Actor, context Context) bool {
		resolved := actor.Resolve(game)
		return resolved.Protected == protected
	}
}
