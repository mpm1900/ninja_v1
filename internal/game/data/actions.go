package data

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/actions"

	"github.com/google/uuid"
)

var ACTIONS map[uuid.UUID]game.Action = map[uuid.UUID]game.Action{
	// SYSTEM
	game.Switch.ID:       game.Switch,
	game.SwitchInIds[1]:  game.SwitchIn(1),
	game.SwitchInIds[2]:  game.SwitchIn(2),
	game.Struggle.ID:     game.Struggle,
	game.CancelSummon.ID: game.CancelSummon,

	// SUMMONS
	actions.ShadowClone.ID:     actions.ShadowClone,
	actions.SummonGamabunta.ID: actions.SummonGamabunta,

	actions.Distraction.ID:     actions.Distraction,
	actions.Taunt.ID:           actions.Taunt,
	actions.BodyReplacement.ID: actions.BodyReplacement,
	actions.Tailwind.ID:        actions.Tailwind,
	actions.Haze.ID:            actions.Haze,
	actions.Flash.ID:           actions.Flash,
	actions.MirageCrow.ID:      actions.MirageCrow,
	actions.Sekiryoku.ID:       actions.Sekiryoku,

	actions.HiddenMist.ID:     actions.HiddenMist,
	actions.Surf.ID:           actions.Surf,
	actions.LuckyStrikes.ID:   actions.LuckyStrikes,
	actions.GreatTreeSpear.ID: actions.GreatTreeSpear,
	actions.C0UltimateArt.ID:  actions.C0UltimateArt,
	actions.C1Bird.ID:         actions.C1Bird,

	actions.GiantRasengan.ID:    actions.GiantRasengan,
	actions.Rasengan.ID:         actions.Rasengan,
	actions.RasenganRecharge.ID: actions.RasenganRecharge,
	actions.Rasenshuriken.ID:    actions.Rasenshuriken,
	actions.FlyingRaijin.ID:     actions.FlyingRaijin,

	actions.HeavyPunch.ID:  actions.HeavyPunch,
	actions.LeafJab.ID:     actions.LeafJab,
	actions.DragonDance.ID: actions.DragonDance,

	// FIRE ATTACKS
	actions.Fireball.ID:              actions.Fireball,
	actions.GreatFireball.ID:         actions.GreatFireball,
	actions.FlameBullet.ID:           actions.FlameBullet,
	actions.PhoenixFlower.ID:         actions.PhoenixFlower,
	actions.DragonFire.ID:            actions.DragonFire,
	actions.SearingMigraine.ID:       actions.SearingMigraine,
	actions.GreatFireAnnihilation.ID: actions.GreatFireAnnihilation,
	actions.Asakujaku.ID:             actions.Asakujaku,

	actions.Chidori.ID:       actions.Chidori,
	actions.ChidoriSpear.ID:  actions.ChidoriSpear,
	actions.ChidoriStream.ID: actions.ChidoriStream,
	actions.Kirin.ID:         actions.Kirin,

	actions.Curse.ID:            actions.Curse,
	actions.Recover.ID:          actions.Recover,
	actions.LeechSeed.ID:        actions.LeechSeed,
	actions.SageMode.ID:         actions.SageMode,
	actions.PowerBoost.ID:       actions.PowerBoost,
	actions.Amaterasu.ID:        actions.Amaterasu,
	actions.Disable.ID:          actions.Disable,
	actions.Coercion.ID:         actions.Coercion,
	actions.BloodPrice.ID:       actions.BloodPrice,
	actions.MindTransfer.ID:     actions.MindTransfer,
	actions.CopyJutsu.ID:        actions.CopyJutsu,
	actions.TempleOfNirvana.ID:  actions.TempleOfNirvana,
	actions.Graft.ID:            actions.Graft,
	actions.InstilFear.ID:       actions.InstilFear,
	actions.HumanBoulder.ID:     actions.HumanBoulder,
	actions.BodyFlicker.ID:      actions.BodyFlicker,
	actions.KamuiCounter.ID:     actions.KamuiCounter,
	actions.KamuiSlash.ID:       actions.KamuiSlash,
	actions.ReaperDeathSeal.ID:  actions.ReaperDeathSeal,
	actions.PerishSong.ID:       actions.PerishSong,
	actions.ShadowPossession.ID: actions.ShadowPossession,
}
