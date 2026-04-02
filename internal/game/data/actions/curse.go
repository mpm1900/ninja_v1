package actions

import (
	"fmt"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var Curse = MakeCurse()

func MakeCurse() game.Action {
	nature := game.NsJashin
	config := game.ActionConfig{
		Name:   "Curse",
		Nature: &nature,
		Jutsu:  game.Fuinjutsu,
	}
	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetActorID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.ActiveFilter, game.AliveFilter),
		ContextValidate: game.TargetLengthFilter(1),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				s, ok := g.GetSource(context)
				if !ok {
					return transactions
				}

				source := s.Resolve(g)
				targets := g.GetTargets(context)

				for _, t := range targets {
					target := t.Resolve(g)
					log := game.AddLogs(
						game.NewLog(fmt.Sprintf("%s was protected.", target.Name)),
					)
					if tx, protected := target.IsProtected(log); protected {
						transactions = append(transactions, tx)
						continue
					}

					hp := source.Stats[game.StatHP]
					amount := hp / 2

					hp_loss := mutations.PureDamageWith(amount, func(a game.Actor) game.Actor {
						if !a.Alive {
							a.Alive = true
							a.Damage = hp - 1
						}
						return a
					})
					damage := mutations.PureDamage(amount)

					sourceTx := game.MakeTransaction(hp_loss, game.WithTargetIDs(context, []uuid.UUID{source.ID}))
					targetTx := game.MakeTransaction(damage, game.WithTargetIDs(context, []uuid.UUID{target.ID}))

					transactions = append(transactions, sourceTx, targetTx)
				}

				return transactions
			},
		},
	}
}
