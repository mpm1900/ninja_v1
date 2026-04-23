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

	actions.SummonAlly.ID:       actions.SummonAlly,
	actions.Distraction.ID:      actions.Distraction,
	actions.Taunt.ID:            actions.Taunt,
	actions.BodyReplacement.ID:  actions.BodyReplacement,
	actions.Barrier.ID:          actions.Barrier,
	actions.Kamui.ID:            actions.Kamui,
	actions.Tailwind.ID:         actions.Tailwind,
	actions.InstilFear.ID:       actions.InstilFear,
	actions.Flash.ID:            actions.Flash,
	actions.MirageCrow.ID:       actions.MirageCrow,
	actions.Haze.ID:             actions.Haze,
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
	actions.KamuiCounter.ID:     actions.KamuiCounter,
	actions.KamuiSlash.ID:       actions.KamuiSlash,
	actions.ReaperDeathSeal.ID:  actions.ReaperDeathSeal,
	actions.PerishSong.ID:       actions.PerishSong,
	actions.ShadowPossession.ID: actions.ShadowPossession,
	actions.Revival.ID:          actions.Revival,
	actions.TeamHeal.ID:         actions.TeamHeal,
	actions.PatternBreak.ID:     actions.PatternBreak,
	actions.ChilliPill.ID:       actions.ChilliPill,

	actions.HiddenMist.ID:     actions.HiddenMist,
	actions.CollidingWave.ID:  actions.CollidingWave,
	actions.WaterDragon.ID:    actions.WaterDragon,
	actions.SharkBomb.ID:      actions.SharkBomb,
	actions.GreatWaterfall.ID: actions.GreatWaterfall,

	actions.GreatTreeSpear.ID: actions.GreatTreeSpear,

	actions.C0UltimateArt.ID: actions.C0UltimateArt,
	actions.C1Bird.ID:        actions.C1Bird,

	actions.FlyingSwallow.ID:    actions.FlyingSwallow,
	actions.Sekiryoku.ID:        actions.Sekiryoku,
	actions.GiantRasengan.ID:    actions.GiantRasengan,
	actions.Rasengan.ID:         actions.Rasengan,
	actions.RasenganRecharge.ID: actions.RasenganRecharge,
	actions.FlyingRaijin.ID:     actions.FlyingRaijin,

	actions.LuckyStrikes.ID:        actions.LuckyStrikes,
	actions.HeavyPunch.ID:          actions.HeavyPunch,
	actions.WhirlwindKick.ID:       actions.WhirlwindKick,
	actions.DragonStance.ID:        actions.DragonStance,
	actions.SwordsStance.ID:        actions.SwordsStance,
	actions.KusariChains.ID:        actions.KusariChains,
	actions.CamelliaDance.ID:       actions.CamelliaDance,
	actions.ClematisDanceFlower.ID: actions.ClematisDanceFlower,
	actions.CherryBlossomImpact.ID: actions.CherryBlossomImpact,

	actions.IronSkin.ID:        actions.IronSkin,
	actions.IronBody.ID:        actions.IronBody,
	actions.Expansion.ID:       actions.Expansion,
	actions.HumanBoulder.ID:    actions.HumanBoulder,
	actions.RockFist.ID:        actions.RockFist,
	actions.EarthDomePrison.ID: actions.EarthDomePrison,

	// FIRE ATTACKS
	actions.Fireball.ID:              actions.Fireball,
	actions.GreatFireball.ID:         actions.GreatFireball,
	actions.FlameBullet.ID:           actions.FlameBullet,
	actions.PhoenixFlower.ID:         actions.PhoenixFlower,
	actions.DragonFire.ID:            actions.DragonFire,
	actions.SearingMigraine.ID:       actions.SearingMigraine,
	actions.GreatFireAnnihilation.ID: actions.GreatFireAnnihilation,
	actions.MajesticFlame.ID:         actions.MajesticFlame,
	actions.Asakujaku.ID:             actions.Asakujaku,
	actions.Firestorm.ID:             actions.Firestorm,
	actions.PunishingFire.ID:         actions.PunishingFire,

	actions.Chidori.ID:        actions.Chidori,
	actions.ChidoriSpear.ID:   actions.ChidoriSpear,
	actions.ChidoriStream.ID:  actions.ChidoriStream,
	actions.Raikiri.ID:        actions.Raikiri,
	actions.Kirin.ID:          actions.Kirin,
	actions.FalseDarkness.ID:  actions.FalseDarkness,
	actions.LightningHound.ID: actions.LightningHound,

	actions.Rasenshuriken.ID:  actions.Rasenshuriken,
	actions.PressureDamage.ID: actions.PressureDamage,
	actions.BodyFlicker.ID:    actions.BodyFlicker,
	actions.VacuumBlast.ID:    actions.VacuumBlast,
	actions.WindSlash.ID:      actions.WindSlash,

	actions.PoisonSting.ID: actions.PoisonSting,

	actions.SnakeStrike.ID: actions.SnakeStrike,
}
