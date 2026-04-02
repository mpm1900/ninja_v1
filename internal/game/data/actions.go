package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var ACTIONS map[uuid.UUID]game.Action = map[uuid.UUID]game.Action{
	game.Switch.ID:      game.Switch,
	game.SwitchInIds[1]: game.SwitchIn(1),
	game.SwitchInIds[2]: game.SwitchIn(2),
	game.SwitchInIds[3]: game.SwitchIn(3),
	game.SwitchInIds[4]: game.SwitchIn(4),
	game.SwitchInIds[5]: game.SwitchIn(5),

	actions.FollowMe.ID:   actions.FollowMe,
	actions.Protect.ID:    actions.Protect,
	actions.Tailwind.ID:   actions.Tailwind,
	actions.Coercion.ID:   actions.Coercion,
	actions.MirageCrow.ID: actions.MirageCrow,

	actions.Surf.ID:         actions.Surf,
	actions.LuckyStrikes.ID: actions.LuckyStrikes,

	actions.Rasengan.ID:         actions.Rasengan,
	actions.RasenganRecharge.ID: actions.RasenganRecharge,

	actions.LeafJab.ID:     actions.LeafJab,
	actions.DragonDance.ID: actions.DragonDance,
	actions.Fireball.ID:    actions.Fireball,
	actions.Chidori.ID:     actions.Chidori,
	actions.Curse.ID:       actions.Curse,
	actions.Recover.ID:     actions.Recover,
	actions.LeechSeed.ID:   actions.LeechSeed,
}
