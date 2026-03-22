package game

import (
	"github.com/google/uuid"
)

type Context struct {
	SourcePlayerID    *uuid.UUID  `json:"source_player_ID"`
	ParentActorID     *uuid.UUID  `json:"parent_actor_ID"`
	SourceActorID     *uuid.UUID  `json:"source_actor_ID"`
	TargetActorIDs    []uuid.UUID `json:"target_actor_IDs"`
	TargetPositionIDs []uuid.UUID `json:"target_position_IDs"`
}
