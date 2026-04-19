package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

var sandstormWeatherID = uuid.MustParse("6dfb1ee5-a0dc-4ee0-8a52-2c062e8d8b3e")
var SandstormWeather = MakeSandstorm()

func MakeSandstorm() game.Modifier {
	sandstormWeather := SetWeather(sandstormWeatherID, game.GameWeatherSandstorm, "Sandstorm")
	// sandstormWeather.ActorMutations = []game.ActorMutation{}
	sandstormWeather.Triggers = append(sandstormWeather.Triggers, game.Trigger{
		ID:         uuid.MustParse("d40e88db-b3cb-4541-889a-9e351b0e44b9"),
		ModifierID: sandstormWeatherID,
		On:         game.OnTurnEnd,
		Check: func(p, g game.Game, context game.Context, tx game.Transaction[game.Modifier]) bool {
			return g.HasWeather(game.GameWeatherSandstorm, context)
		},
		ActionMutation: game.ActionMutation{
			Priority: game.ActionPriorityDefault,
			Filter:   game.TrueGameFilter,
			Delta: func(p game.Game, g game.Game, context game.Context) []game.Transaction[game.GameMutation] {
				mut := game.RatioDamage(0.0625)
				mut_ctx := context
				mut_ctx.TargetActorIDs = []uuid.UUID{}
				for _, target := range g.GetActiveActors() {
					mut_ctx.TargetActorIDs = append(mut_ctx.TargetActorIDs, target.ID)
				}
				return []game.Transaction[game.GameMutation]{
					game.MakeTransaction(mut, mut_ctx),
				}
			},
		},
	})
	return sandstormWeather
}

var sandAuraID = uuid.MustParse("b16bfb0c-8131-4522-8f97-7c5775e7df05")

var SandAuraTrigger game.Trigger = game.Trigger{
	ID:         uuid.New(),
	ModifierID: sandAuraID,
	On:         game.OnActorEnter,
	Check:      game.Match__SourceActor_SourceActor,
	ActionMutation: game.ActionMutation{
		Priority: game.ActionPriorityDefault,
		Filter:   game.TrueGameFilter,
		Delta: func(p game.Game, g game.Game, context game.Context) []game.GameTransaction {
			transactions := []game.GameTransaction{}

			state, _ := g.GetState(context)
			if state.Weather == game.GameWeatherSandstorm {
				return transactions
			}

			mod := SandstormWeather
			mod.Duration = 4
			mutation := mutations.AddModifiers(false, mod)
			transaction := game.MakeTransaction(mutation, game.NewContext())
			transactions = append(transactions, transaction)

			return transactions
		},
	},
}

var SandAura = game.Modifier{
	ID:          sandAuraID,
	GroupID:     &sandAuraID,
	Name:        "Sand Aura",
	Description: "On enter: start sandstorm.",
	Show:        true,
	Duration:    game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
		game.NewNoopSource(&sandAuraID),
	},
	Triggers: []game.Trigger{
		SandAuraTrigger,
	},
}
