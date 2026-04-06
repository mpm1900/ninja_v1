package actions

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var LeechSeed = MakeLeechSeed()

func MakeLeechSeed() game.Action {
	nature := game.NsYang
	targetCount := 1
	chakraCost := 30
	config := game.ActionConfig{
		Name:        "Leech Seed",
		Nature:      &nature,
		TargetCount: &targetCount,
		Cost:        &chakraCost,
		Jutsu:       game.Senjutsu,
	}

	return game.Action{
		ID:              uuid.New(),
		Config:          config,
		TargetType:      game.TargetPositionID,
		TargetPredicate: game.ComposeAF(game.ActiveFilter),
		ContextValidate: game.PositionsLengthFilter(*config.TargetCount),
		Cost:            mutations.UseStaminaSource(chakraCost),
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.SourceIsAlive,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
				transactions := []game.GameTransaction{}

				source, ok := g.GetSource(context)
				if !ok || source.PositionID == nil {
					return transactions
				}

				targets := g.GetTargets(context)
				for _, target := range targets {
					mut_ctx := game.Context{
						SourcePlayerID:    &target.PlayerID,
						SourceActorID:     &source.ID,
						ParentActorID:     &target.ID,
						TargetPositionIDs: []uuid.UUID{*source.PositionID},
					}
					mutation := mutations.AddModifiers(LeechSeedModifier)
					transaction := game.MakeTransaction(mutation, mut_ctx)
					transactions = append(transactions, transaction)
				}

				return transactions
			},
		},
	}
}

var leechSeedModifierID = uuid.New()

var LeechSeedTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: leechSeedModifierID,
	On:         game.OnTurnEnd,
	Check: func(p game.Game, g game.Game, ctx game.Context, t game.Transaction[game.Modifier]) bool {
		return true
	},
	ActionMutation: game.ActionMutation{
		Priority: 0,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
			transactions := []game.GameTransaction{}
			parent, ok := g.GetParent(context)
			if !ok {
				return transactions
			}

			resolved_parent := parent.Resolve(g)
			targets := g.GetTargets(context)
			for _, target := range targets {
				ratio := 0.125
				hp_loss := game.Round(float64(resolved_parent.Stats[game.StatHP]) * ratio)
				hp_loss_ctx := context
				hp_loss_ctx.TargetActorIDs = []uuid.UUID{resolved_parent.ID}
				hp_loss_ctx.TargetPositionIDs = []uuid.UUID{}
				hp_loss_mut := mutations.PureDamage(hp_loss, false)
				hp_loss_tx := game.MakeTransaction(hp_loss_mut, hp_loss_ctx)

				heal_mut := mutations.PureHeal(hp_loss)
				heal_ctx := context
				heal_ctx.TargetActorIDs = []uuid.UUID{target.ID}
				heal_ctx.TargetPositionIDs = []uuid.UUID{}
				heal_tx := game.MakeTransaction(heal_mut, heal_ctx)

				transactions = append(transactions, hp_loss_tx, heal_tx)
			}

			return transactions
		},
	},
}

var LeechSeedModifier game.Modifier = game.Modifier{
	ID:       leechSeedModifierID,
	GroupID:  &leechSeedModifierID,
	Name:     "Seeded",
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopParent(&leechSeedModifierID),
	},
	Triggers: []game.Trigger{
		LeechSeedTrigger,
	},
}
