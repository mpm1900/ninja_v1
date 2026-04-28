package actors

import (
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var GlobalActions = []uuid.UUID{
	actions.BodyReplacement.ID,
	actions.Rest.ID,
	actions.ShadowClone.ID,
}
