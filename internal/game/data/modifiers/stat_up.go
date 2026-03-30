package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewStageDelta(
	stat game.BaseStat,
	groupID uuid.UUID,
	filter func(input game.Actor, context game.Context) bool,
	priority int,
	delta int,
) game.Modifier {
	mut := game.MakeModifierMutation(
		&groupID,
		priority,
		filter,
		func(actor game.Actor, context game.Context) game.Actor {
			actor.Stages[stat] = actor.Stages[stat] + delta
			return actor
		},
	)

	return game.Modifier{
		ID:      uuid.New(),
		GroupID: groupID,
		Mutations: []game.ModifierMutation{
			mut,
		},
	}
}

func MakeStatUpSource(stat game.BaseStat, name string, groupID uuid.UUID) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ComposeAF(game.ActiveFilter, game.SourceFilter), game.MutPriorityDefault, 1)
	modifier.Name = name
	return modifier
}

func MakeStatUpAll(stat game.BaseStat, name string, groupID uuid.UUID) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.ActiveFilter, game.MutPriorityDefault, 1)
	modifier.Name = name
	return modifier
}

var AttackUpID = uuid.New()
var JutsuUpID = uuid.New()
var SpeedUpID = uuid.New()

var AttackUpSource = MakeStatUpSource(game.StatAttack, "Attack Up", AttackUpID)
var JutsuUpSource = MakeStatUpSource(game.BaseStat(game.ChakraAttack), "Chakra Attack Up", JutsuUpID)
var SpeedUpSource = MakeStatUpSource(game.StatSpeed, "Speed Up", SpeedUpID)
var SpeedUpAll = MakeStatUpAll(game.StatSpeed, "Speed Up", SpeedUpID)
