package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewStageDelta(
	stat game.BaseStat,
	groupID uuid.UUID,
	filter func(input game.Actor, context *game.Context) bool,
	priority int,
	delta int,
) game.Modifier {
	mut := game.MakeModifierMutation(
		&groupID,
		priority,
		filter,
		func(actor game.Actor, context *game.Context) game.Actor {
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
	modifier := NewStageDelta(stat, groupID, game.SourceFilter, game.PriorityDefault, 1)
	modifier.Name = name
	return modifier
}

func MakeStatUpAll(stat game.BaseStat, name string, groupID uuid.UUID) game.Modifier {
	modifier := NewStageDelta(stat, groupID, game.AllFilter, game.PriorityDefault, 1)
	modifier.Name = name
	return modifier
}

var GenjutsuUpID = uuid.New()
var SpeedUpId = uuid.New()
var TaijutsuUpID = uuid.New()

var GenjutsuUpSource = MakeStatUpSource(game.StatGenjutsu, "Genjutsu Up", GenjutsuUpID)
var SpeedUpSource = MakeStatUpSource(game.StatSpeed, "Speed Up", SpeedUpId)
var SpeedUpAll = MakeStatUpAll(game.StatSpeed, "Speed Up", SpeedUpId)
var TaijutsuUpSource = MakeStatUpSource(game.StatTaijutsu, "Taijutsu Up", TaijutsuUpID)
