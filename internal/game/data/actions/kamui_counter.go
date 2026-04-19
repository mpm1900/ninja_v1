package actions

import (
	"fmt"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var KamuiCounter = MakeBlitz()

func MakeBlitz() game.Action {
	config := game.ActionConfig{
		Name:        "Kamui: Counter",
		Description: "+1 priority. Fails unless the target has a pending attack.",
		Accuracy:    game.Ptr(100),
		Power:       game.Ptr(70),
		Stat:        game.Ptr(game.StatAttack),
		Nature:      game.Ptr(game.NsYin),
		Cost:        game.Ptr(0),
		TargetCount: game.Ptr(1),
		Jutsu:       game.Taijutsu,
		CritChance:  game.Ptr(5),
		CritMod:     1.5,
	}
	return game.Action{
		ID:              uuid.MustParse("09b013e2-436d-4bd4-bf0a-9b2827b9c131"),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.OtherFilter, game.TargetableFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(*config.Cost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityP1,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				// loop through every taerget and every action to see if they have a pending attack, if not, this attack fails
				for _, target := range g.GetTargets(context) {
					var found_action *game.Transaction[game.Action] = nil
					for _, action := range g.Actions {
						if action.Context.SourceActorID != nil && *action.Context.SourceActorID == target.ID {
							// attacking moves only
							if action.Mutation.Config.Power != nil {
								found_action = &action
								break
							}

						}
					}

					if found_action == nil {
						// then target has no action
						log := game.NewLog(fmt.Sprintf("%s failed.", config.Name))
						log_mut := game.AddLogs(log)
						log_tx := game.MakeTransaction(log_mut, game.NewContext())
						transactions = append(transactions, log_tx)
						return transactions
					}
				}

				conf := game.GetActiveActionConfig(g, config)
				crit_result := game.MakeCriticalCheck(conf)
				damages := mutations.NewDamage(conf, game.NewDamageConfig(crit_result.Ratio, game.RandomDamageFactor()))
				transactions = append(
					transactions,
					mutations.MakeDamageTransactions(context, damages)...,
				)

				return transactions
			},
		},
	}
}
