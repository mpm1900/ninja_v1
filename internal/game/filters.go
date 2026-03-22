package game

func AllFilter(actor Actor, context *Context) bool {
	return true
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
