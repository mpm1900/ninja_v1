package game

import (
	"slices"
)

/**
 * Context Filters
 */

type ContextFilter func(Context) bool

func ComposeCF(filters ...ContextFilter) ContextFilter {
	return func(context Context) bool {
		for _, filter := range filters {
			if !filter(context) {
				return false
			}
		}

		return true
	}
}

func TargetLengthFilter(length int) func(Context) bool {
	return func(context Context) bool {
		return len(context.TargetActorIDs) == length
	}
}
func PositionsLengthFilter(length int) func(Context) bool {
	return func(context Context) bool {
		return len(context.TargetPositionIDs) == length
	}
}

/**
 * Actor Filters
 */

type ActorFilter func(Actor, Context) bool

func ComposeAF(filters ...ActorFilter) ActorFilter {
	return func(actor Actor, context Context) bool {
		for _, filter := range filters {
			if !filter(actor, context) {
				return false
			}
		}

		return true
	}
}
func AllFilter(actor Actor, context Context) bool {
	return true
}
func NoneFilter(actor Actor, context Context) bool {
	return false
}
func OtherFilter(actor Actor, context Context) bool {
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
func AliveFilter(actor Actor, context Context) bool {
	return actor.Alive
}
func ActiveFilter(actor Actor, context Context) bool {
	return actor.PositionID != nil
}
func InactiveFilter(actor Actor, context Context) bool {
	return actor.PositionID == nil
}
func SourceFilter(actor Actor, context Context) bool {
	if context.SourceActorID == nil {
		return false
	}
	if !ActiveFilter(actor, context) {
		return false
	}
	return actor.ID == *context.SourceActorID
}
func ParentFilter(actor Actor, context Context) bool {
	if context.ParentActorID == nil {
		return false
	}
	if !ActiveFilter(actor, context) {
		return false
	}
	return actor.ID == *context.ParentActorID
}
func TargetFilter(actor Actor, context Context) bool {
	if slices.Contains(context.TargetActorIDs, actor.ID) {
		return true
	}
	if actor.PositionID != nil && slices.Contains(context.TargetPositionIDs, *actor.PositionID) {
		return true
	}

	return false
}
func TeamFilter(actor Actor, context Context) bool {
	if context.SourcePlayerID == nil {
		return false
	}
	return actor.PlayerID == *context.SourcePlayerID
}
func OtherTeamFilter(actor Actor, context Context) bool {
	if context.SourcePlayerID == nil {
		return false
	}
	return actor.PlayerID != *context.SourcePlayerID
}

/**
 * RESOLVED FILTERS
 * These filters required modifiers to be resolved to check things like
 * protected, chakra amount, and things that can change with modifers
 */
func HasChakraFilter(game Game, amount int) func(Actor, Context) bool {
	return func(actor Actor, context Context) bool {
		resolved := actor.Resolve(game)
		return resolved.HasChakra(amount)
	}
}
func IsProtectedFilter(game Game, protected bool) func(Actor, Context) bool {
	return func(actor Actor, context Context) bool {
		resolved := actor.Resolve(game)
		return resolved.Protected == protected
	}
}

/**
 * Game Filters
 */

type GameFilter func(Game, Context) bool

func ComposeGF(filters ...GameFilter) GameFilter {
	return func(game Game, context Context) bool {
		for _, filter := range filters {
			if !filter(game, context) {
				return false
			}
		}

		return true
	}
}

func AllGameFilter(game Game, context Context) bool {
	return true
}
func SourceIsAlive(game Game, context Context) bool {
	source, ok := game.GetSource(context)
	if !ok {
		return false
	}

	return source.Alive
}
func SourceIsActionOnCooldown(g Game, context Context) bool {
	if context.ActionID == nil {
		return false
	}

	source, ok := g.GetSource(context)
	if !ok {
		return false
	}

	cooldown, ok := source.ActionCooldowns[*context.ActionID]
	if !ok {
		return true
	}

	return cooldown == 0
}
func SourceHasActiveTurns(turns int) func(Game, Context) bool {
	return func(g Game, context Context) bool {
		if context.ActionID == nil {
			return false
		}

		source, ok := g.GetSource(context)
		if !ok {
			return false
		}

		return source.ActiveTurns == turns
	}
}
func TargetsIsOneAlive(game Game, context Context) bool {
	targets := game.GetTargets(context)
	for _, target := range targets {
		if target.Alive {
			return true
		}
	}
	return false
}

func Match__TargetActor_SourceActor(game Game, context Context, modifier_tx Transaction[Modifier]) bool {
	targets := game.GetTargets(context)
	if len(targets) == 0 || modifier_tx.Context.SourceActorID == nil {
		return false
	}

	for _, t := range targets {
		if t.ID == *modifier_tx.Context.SourceActorID {
			return true
		}
	}
	return false
}
func Match__SourceActor_SourceActor(game Game, context Context, modifier_tx Transaction[Modifier]) bool {
	if context.SourceActorID == nil || modifier_tx.Context.SourceActorID == nil {
		return false
	}

	return *context.SourceActorID == *modifier_tx.Context.SourceActorID
}
