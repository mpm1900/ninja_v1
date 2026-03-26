package game

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
	return actor.ID != *context.SourceActorID
}
func AliveFilter(actor Actor, context Context) bool {
	return actor.State.Alive
}
func ActiveFilter(actor Actor, context Context) bool {
	return actor.State.PositionID != nil
}
func InactiveFilter(actor Actor, context Context) bool {
	return actor.State.PositionID == nil
}
func SourceFilter(actor Actor, context Context) bool {
	if !ActiveFilter(actor, context) {
		return false
	}
	return actor.ID == *context.SourceActorID
}
func TeamFilter(actor Actor, context Context) bool {
	return actor.PlayerID == *context.SourcePlayerID
}
func OtherTeamFilter(actor Actor, context Context) bool {
	return actor.PlayerID != *context.SourcePlayerID
}

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

func MatchSourceActorIDTrigger(game Game, context Context, modifier_tx Transaction[Modifier]) bool {
	if context.SourceActorID == nil || modifier_tx.Context.SourceActorID == nil {
		return false
	}

	return *context.SourceActorID == *modifier_tx.Context.SourceActorID
}

func Match__TargetActor_SourceActor(game Game, context Context, modifier_tx Transaction[Modifier]) bool {
	if len(context.TargetActorIDs) == 0 || modifier_tx.Context.SourceActorID == nil {
		return false
	}

	return context.TargetActorIDs[0] == *modifier_tx.Context.SourceActorID
}
