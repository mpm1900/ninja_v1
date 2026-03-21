package game

import (
	"github.com/google/uuid"
)

type Context struct {
	SourcePlayerID    *uuid.UUID
	ParentActorID     *uuid.UUID
	SourceActorID     *uuid.UUID
	TargetActorIDs    []uuid.UUID
	TargetPositionIDs []uuid.UUID
}
