package actors

import (
	"ninja_v1/internal/game/data/actions"
	"slices"

	"github.com/google/uuid"
)

var GlobalActions = []uuid.UUID{
	actions.BodyReplacement.ID,
	actions.LightningKunai.ID,
	actions.Fireball.ID,
	actions.GalePalm.ID,
	actions.Rest.ID,
	actions.ShadowClone.ID,
	actions.StoneBullet.ID,
	actions.WaterBullet.ID,
}

func GlobalActionsExcept(ids ...uuid.UUID) []uuid.UUID {
	global := slices.Clone(GlobalActions)
	return slices.DeleteFunc(global, func(id uuid.UUID) bool {
		return slices.Contains(ids, id)
	})
}
