package modifiers

import (
	"math"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewStageDelta(
	stat game.ActorStat,
	groupID *uuid.UUID,
	filter func(input game.Actor, context game.Context) bool,
	priority int,
	delta int,
) game.Modifier {
	mut := game.MakeActorMutation(
		groupID,
		priority,
		filter,
		func(g game.Game, actor game.Actor, context game.Context) game.Actor {
			actor.Stages[stat] = actor.Stages[stat] + delta
			return actor
		},
	)

	return game.Modifier{
		ID:       uuid.New(),
		GroupID:  groupID,
		Duration: game.ModifierDurationInf,
		Mutations: []game.ModifierMutation{
			mut,
		},
	}
}

func NewStatMult(
	stat game.ActorStat,
	groupID *uuid.UUID,
	filter func(input game.Actor, context game.Context) bool,
	priority int,
	mult float64,
) game.Modifier {
	mut := game.MakeActorMutation(
		groupID,
		priority,
		filter,
		func(g game.Game, actor game.Actor, context game.Context) game.Actor {
			actor.Stats[stat] = int(math.Floor(float64(actor.Stats[stat]) * mult))
			return actor
		},
	)

	return game.Modifier{
		ID:       uuid.New(),
		GroupID:  groupID,
		Duration: game.ModifierDurationInf,
		Mutations: []game.ModifierMutation{
			mut,
		},
	}
}

func MakeStatDeltaSource(stat game.ActorStat, name string, groupID *uuid.UUID, delta int) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ComposeAF(game.ActiveFilter, game.SourceFilter), game.MutPriorityDefault, delta)
	modifier.Name = name
	return modifier
}

func MakeStatDeltaTeam(stat game.ActorStat, name string, groupID *uuid.UUID, delta int) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ComposeAF(game.ActiveFilter, game.TeamFilter), game.MutPriorityDefault, delta)
	modifier.Name = name
	return modifier
}

func MakeStatDeltaAll(stat game.ActorStat, name string, groupID *uuid.UUID, delta int) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ActiveFilter, game.MutPriorityDefault, delta)
	modifier.Name = name
	return modifier
}

func MakeStatMultTeam(stat game.ActorStat, name string, groupID *uuid.UUID, mult float64) game.Modifier {
	modifier := NewStatMult(stat, groupID, game.ComposeAF(game.ActiveFilter, game.TeamFilter), game.MutPriorityPostStagedStats, mult)
	modifier.Name = name
	return modifier
}

var AttackUpSource = MakeStatDeltaSource(game.StatAttack, "Attack Up", nil, 1)
var AttackDownSource = MakeStatDeltaSource(game.StatAttack, "Attack Down", nil, -1)
var ChakraAttackUpSource = MakeStatDeltaSource(game.StatChakraAttack, "Chakra Attack Up", nil, 1)
var ChakraAttackDownSource = MakeStatDeltaSource(game.StatChakraAttack, "Chakra Attack Down", nil, -1)
var ChakraAttackDown2Source = MakeStatDeltaSource(game.StatChakraAttack, "Chakra Attack Down (2)", nil, -2)
var SpeedUpSource = MakeStatDeltaSource(game.StatSpeed, "Speed Up", nil, 1)
var SpeedUpTeam = MakeStatDeltaTeam(game.StatSpeed, "Speed Up", nil, 1)
var SpeedUpAll = MakeStatDeltaAll(game.StatSpeed, "Speed Up", nil, 1)

// NAMED STAT UPS
var TailwindID = uuid.New()
var Tailwind = MakeStatMultTeam(game.StatSpeed, "Tailwind", &TailwindID, 2.0)

// HAZE
var hazeID = uuid.New()
var Haze game.Modifier = game.Modifier{
	ID:       hazeID,
	GroupID:  &hazeID,
	Name:     "Haze",
	Duration: game.ModifierDurationInf,
	Mutations: []game.ModifierMutation{
		game.MakeActorMutation(
			&hazeID,
			game.MutPrioritySet,
			game.ActiveFilter,
			func(g game.Game, actor game.Actor, context game.Context) game.Actor {
				actor.Stages[game.StatAccuracy] = 0
				actor.Stages[game.StatAttack] = 0
				actor.Stages[game.StatChakraAttack] = 0
				actor.Stages[game.StatChakraDefense] = 0
				actor.Stages[game.StatDefense] = 0
				actor.Stages[game.StatEvasion] = 0
				actor.Stages[game.StatHP] = 0
				actor.Stages[game.StatSpeed] = 0
				actor.Stages[game.StatStamina] = 0

				return actor
			},
		),
	},
}
