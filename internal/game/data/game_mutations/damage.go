package mutations

import (
	"ninja_v1/internal/game"
)

func ApplyDamage(g game.Game, target game.ResolvedActor, damage int) {
	g.UpdateActor(target.ID, func(a game.Actor) game.Actor {
		a.Damage += damage
		a.Alive = target.Stats[game.StatHP] > a.Damage
		return a
	})
}

func NewDamage(stat game.AttackStat, power int, nature *game.NatureSet) *game.GameMutation {
	return &game.GameMutation{
		Delta: func(g game.Game, context *game.Context) game.Game {
			ok, s := g.GetActor(func(a game.Actor) bool {
				return a.ID == *context.SourceActorID
			})

			if !ok {
				return g
			}

			targets := game.GetTargets(g, *context)
			source := game.ResolveActor(s, g)
			for _, t := range targets {
				target := game.ResolveActor(t, g)
				damages := game.GetDamage(
					source,
					[]game.ResolvedActor{target},
					stat,
					power,
					nature,
				)

				for _, damage := range damages {
					ApplyDamage(g, target, damage)
				}
			}

			return g
		},
	}
}
