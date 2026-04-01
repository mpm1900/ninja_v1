package modifiers

import (
	"math"
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewStageDelta(
	stat game.ActorStat,
	groupID uuid.UUID,
	filter func(input game.Actor, context game.Context) bool,
	priority int,
	delta int,
) game.Modifier {
	mut := game.MakeActorMutation(
		&groupID,
		priority,
		filter,
		func(actor game.Actor, context game.Context) game.Actor {
			actor.Stages[stat] = actor.Stages[stat] + delta
			return actor
		},
	)

	return game.Modifier{
		ID:       uuid.New(),
		GroupID:  groupID,
		Duration: game.ModifierDurationInf,
		Mutations: []game.ActorMutation{
			mut,
		},
	}
}

func NewStatMult(
	stat game.ActorStat,
	groupID uuid.UUID,
	filter func(input game.Actor, context game.Context) bool,
	priority int,
	mult float64,
) game.Modifier {
	mut := game.MakeActorMutation(
		&groupID,
		priority,
		filter,
		func(actor game.Actor, context game.Context) game.Actor {
			actor.Stats[stat] = int(math.Floor(float64(actor.Stats[stat]) * mult))
			return actor
		},
	)

	return game.Modifier{
		ID:       uuid.New(),
		GroupID:  groupID,
		Duration: game.ModifierDurationInf,
		Mutations: []game.ActorMutation{
			mut,
		},
	}
}

func MakeStatDeltaSource(stat game.ActorStat, name string, groupID uuid.UUID, delta int) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ComposeAF(game.ActiveFilter, game.SourceFilter), game.MutPriorityDefault, delta)
	modifier.Name = name
	return modifier
}

func MakeStatDeltaTeam(stat game.ActorStat, name string, groupID uuid.UUID, delta int) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ComposeAF(game.ActiveFilter, game.TeamFilter), game.MutPriorityDefault, delta)
	modifier.Name = name
	return modifier
}

func MakeStatDeltaAll(stat game.ActorStat, name string, groupID uuid.UUID, delta int) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ActiveFilter, game.MutPriorityDefault, delta)
	modifier.Name = name
	return modifier
}

func MakeStatMultTeam(stat game.ActorStat, name string, groupID uuid.UUID, mult float64) game.Modifier {
	modifier := NewStatMult(stat, groupID, game.ComposeAF(game.ActiveFilter, game.TeamFilter), game.MutPriorityPostStagedStats, mult)
	modifier.Name = name
	return modifier
}

var AttackUpID = uuid.New()
var AttackDownID = uuid.New()
var JutsuUpID = uuid.New()
var SpeedUpID = uuid.New()

var AttackUpSource = MakeStatDeltaSource(game.StatAttack, "Attack Up", AttackUpID, 1)
var AttackDownSource = MakeStatDeltaSource(game.StatAttack, "Attack Down", AttackDownID, -1)
var JutsuUpSource = MakeStatDeltaSource(game.ActorStat(game.ChakraAttack), "Chakra Attack Up", JutsuUpID, 1)
var SpeedUpSource = MakeStatDeltaSource(game.StatSpeed, "Speed Up", SpeedUpID, 1)
var SpeedUpTeam = MakeStatDeltaTeam(game.StatSpeed, "Speed Up", SpeedUpID, 1)
var SpeedUpAll = MakeStatDeltaAll(game.StatSpeed, "Speed Up", SpeedUpID, 1)

// NAMED STAT UPS
var TailwindID = uuid.New()
var Tailwind = MakeStatMultTeam(game.StatSpeed, "Tailwind", TailwindID, 2.0)
