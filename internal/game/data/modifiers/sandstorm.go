package modifiers

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data/mutations"

	"github.com/google/uuid"
)

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

			filter := FilterWeather()
			transactions = append(transactions, filter)

			mod := SandstormWeather()
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
	Icon:        "sand_aura",
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
