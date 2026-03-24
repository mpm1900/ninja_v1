package game

// CONTEXT FILTERS
// ==================
func TargetLengthFilter(length int) func(*Context) bool {
	return func(context *Context) bool {
		return len(context.TargetActorIDs) == length
	}
}

// ACTOR FILTERS
// ================
func AllFilter(actor Actor, context *Context) bool {
	return true
}
func NoneFilter(actor Actor, context *Context) bool {
	return false
}
func OtherFilter(actor Actor, context *Context) bool {
	return actor.ID != *context.SourceActorID
}

func ActiveFilter(actor Actor, context *Context) bool {
	return actor.Active
}
func SourceFilter(actor Actor, context *Context) bool {
	if !ActiveFilter(actor, context) {
		return false
	}
	return actor.ID == *context.SourceActorID
}

func TeamFilter(actor Actor, context *Context) bool {
	if !ActiveFilter(actor, context) {
		return false
	}
	return actor.PlayerID == *context.SourcePlayerID
}

// GAME FILTERS
// ===============
func AllGameFilter(game Game, context *Context) bool {
	return true
}
