package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewStageDelta(
	stat game.ActorStat,
	groupID *uuid.UUID,
	filter func(game.Game, game.Actor, game.Context) bool,
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
		ActorMutations: []game.ActorMutation{
			mut,
		},
	}
}

func NewStatMult(
	stat game.ActorStat,
	groupID *uuid.UUID,
	filter func(game.Game, game.Actor, game.Context) bool,
	priority int,
	mult float64,
) game.Modifier {
	mut := game.MakeActorMutation(
		groupID,
		priority,
		filter,
		func(g game.Game, actor game.Actor, context game.Context) game.Actor {
			actor.Stats[stat] = game.Round(float64(actor.Stats[stat]) * mult)
			return actor
		},
	)

	return game.Modifier{
		ID:       uuid.New(),
		GroupID:  groupID,
		Duration: game.ModifierDurationInf,
		ActorMutations: []game.ActorMutation{
			mut,
		},
	}
}

func MakeStatDeltaSource(stat game.ActorStat, groupID *uuid.UUID, delta int) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ComposeAF(game.ActiveFilter, game.SourceFilter), game.MutPriorityDefault, delta)
	return modifier
}
func MakeStatDeltaSourceWithName(stat game.ActorStat, groupID *uuid.UUID, delta int, name string) game.Modifier {
	modifier := MakeStatDeltaSource(stat, groupID, delta)
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

func MakeStatMultTeam(stat game.ActorStat, name string, groupID *uuid.UUID, mult float64, priority int) game.Modifier {
	modifier := NewStatMult(stat, groupID, game.ComposeAF(game.ActiveFilter, game.TeamFilter), priority, mult)
	modifier.Name = name
	return modifier
}

var AttackUpID = uuid.New()
var AttackDownID = uuid.New()
var AttackUpSource = MakeStatDeltaSource(game.StatAttack, &AttackUpID, 1)
var AttackUp2Source = MakeStatDeltaSource(game.StatAttack, &AttackUpID, 2)
var AttackDownSource = MakeStatDeltaSource(game.StatAttack, &AttackDownID, -1)
var AttackDown2Source = MakeStatDeltaSource(game.StatAttack, &AttackDownID, -2)
var DefenseUpID = uuid.New()
var DefenseDownID = uuid.New()
var DefenseUpSource = MakeStatDeltaSource(game.StatDefense, &DefenseUpID, 1)
var DefenseUp2Source = MakeStatDeltaSource(game.StatDefense, &DefenseUpID, 2)
var DefenseDownSource = MakeStatDeltaSource(game.StatDefense, &DefenseDownID, -1)
var DefenseDown2Source = MakeStatDeltaSource(game.StatDefense, &DefenseDownID, -2)
var ChakraAttackUpID = uuid.New()
var ChakraAttackDownID = uuid.New()
var ChakraAttackUpSource = MakeStatDeltaSource(game.StatChakraAttack, &ChakraAttackUpID, 1)
var ChakraAttackUp2Source = MakeStatDeltaSource(game.StatChakraAttack, &ChakraAttackUpID, 2)
var ChakraAttackDownSource = MakeStatDeltaSource(game.StatChakraAttack, &ChakraAttackDownID, -1)
var ChakraAttackDown2Source = MakeStatDeltaSource(game.StatChakraAttack, &ChakraAttackDownID, -2)
var ChakraDefenseUpID = uuid.New()
var ChakraDefenseDownID = uuid.New()
var ChakraDefenseUpSource = MakeStatDeltaSource(game.StatChakraDefense, &ChakraDefenseUpID, 1)
var ChakraDefenseUp2Source = MakeStatDeltaSource(game.StatChakraDefense, &ChakraDefenseUpID, 2)
var ChakraDefenseDownSource = MakeStatDeltaSource(game.StatChakraDefense, &ChakraDefenseDownID, -1)
var ChakraDefenseDown2Source = MakeStatDeltaSource(game.StatChakraDefense, &ChakraDefenseDownID, -2)
var SpeedUpID = uuid.New()
var SpeedDownID = uuid.New()
var SpeedUpSource = MakeStatDeltaSource(game.StatSpeed, &SpeedUpID, 1)
var SpeedUp2Source = MakeStatDeltaSource(game.StatSpeed, &SpeedUpID, 2)
var SpeedDownSource = MakeStatDeltaSource(game.StatSpeed, &SpeedDownID, -1)
var SpeedDown2Source = MakeStatDeltaSource(game.StatSpeed, &SpeedDownID, -2)

var EvasionUpID = uuid.New()
var EvasionUpSource = MakeStatDeltaSourceWithName(game.StatEvasion, &EvasionUpID, 1, "Evasion Up")
var AccuracyUpID = uuid.New()
var AccuracyUpSource = MakeStatDeltaSourceWithName(game.StatAccuracy, &AccuracyUpID, 1, "Accuracy Up")

// NAMED STAT UPS
var TailwindID = uuid.New()
var Tailwind = MakeStatMultTeam(game.StatSpeed, "Tailwind", &TailwindID, 2.0, game.MutPriorityPostStagedStats)

var ToadSongID = uuid.New()
var ToadSong = MakeStatMultTeam(game.StatSpeed, "Toad Song", &ToadSongID, -1, game.MutPriorityPostSet)

// HAZE
var hazeID = uuid.New()
var Haze game.Modifier = game.Modifier{
	ID:       hazeID,
	GroupID:  &hazeID,
	Name:     "Haze",
	Duration: game.ModifierDurationInf,
	ActorMutations: []game.ActorMutation{
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
