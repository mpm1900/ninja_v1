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

	actions.Substitution.ID:    actions.Substitution,
	actions.SummonGamabunta.ID: actions.SummonGamabunta,

	actions.Distraction.ID: actions.Distraction,
	actions.Taunt.ID:       actions.Taunt,
	actions.Protect.ID:     actions.Protect,
	actions.Tailwind.ID:    actions.Tailwind,
	actions.Haze.ID:        actions.Haze,
	actions.Flash.ID:       actions.Flash,
	actions.MirageCrow.ID:  actions.MirageCrow,
	actions.Sekiryoku.ID:   actions.Sekiryoku,

	actions.Surf.ID:            actions.Surf,
	actions.LuckyStrikes.ID:    actions.LuckyStrikes,
	actions.GreatTreeSpear.ID:  actions.GreatTreeSpear,
	actions.C0UltimateArt.ID:   actions.C0UltimateArt,
	actions.C1Bird.ID:          actions.C1Bird,
	actions.SearingMigraine.ID: actions.SearingMigraine,

	actions.GiantRasengan.ID:    actions.GiantRasengan,
	actions.Rasengan.ID:         actions.Rasengan,
	actions.RasenganRecharge.ID: actions.RasenganRecharge,
	actions.Rasenshuriken.ID:    actions.Rasenshuriken,

	actions.LeafJab.ID:               actions.LeafJab,
	actions.DragonDance.ID:           actions.DragonDance,
	actions.Fireball.ID:              actions.Fireball,
	actions.Chidori.ID:               actions.Chidori,
	actions.Curse.ID:                 actions.Curse,
	actions.Recover.ID:               actions.Recover,
	actions.LeechSeed.ID:             actions.LeechSeed,
	actions.ToadSong.ID:              actions.ToadSong,
	actions.PowerBoost.ID:            actions.PowerBoost,
	actions.Amaterasu.ID:             actions.Amaterasu,
	actions.Disable.ID:               actions.Disable,
	actions.Coercion.ID:              actions.Coercion,
	actions.BloodPrice.ID:            actions.BloodPrice,
	actions.MindTransfer.ID:          actions.MindTransfer,
	actions.CopyJutsu.ID:             actions.CopyJutsu,
	actions.TempleOfNirvana.ID:       actions.TempleOfNirvana,
	actions.Graft.ID:                 actions.Graft,
	actions.InstilFear.ID:            actions.InstilFear,
	actions.HumanBoulder.ID:          actions.HumanBoulder,
	actions.BodyFlicker.ID:           actions.BodyFlicker,
	actions.KamuiCounter.ID:          actions.KamuiCounter,
	actions.KamuiSlash.ID:            actions.KamuiSlash,
	actions.ReaperDeathSeal.ID:       actions.ReaperDeathSeal,
	actions.PerishSong.ID:            actions.PerishSong,
	actions.ShadowPossession.ID:      actions.ShadowPossession,
	actions.GreatFireAnnihilation.ID: actions.GreatFireAnnihilation,
}
