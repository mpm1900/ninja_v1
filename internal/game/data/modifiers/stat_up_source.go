package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewStatUpSource(stat game.BaseStat, name string, modifierID uuid.UUID) game.Modifier {
	mut := game.MakeModifierMutation(
		&modifierID,
		game.PriorityDefault,
		game.SourceFilter,
		func(actor game.Actor, context *game.Context) game.Actor {
			actor.Stages[stat] = actor.Stages[stat] + 1
			return actor
		},
	)

	return game.Modifier{
		ID:   modifierID,
		Name: name,
		Mutations: []game.ModifierMutation{
			mut,
		},
	}
}

var GenjutsuUpSource = NewStatUpSource(game.StatGenjutsu, "Genjutsu Up", uuid.New())
var SpeedUp = NewStatUpSource(game.StatSpeed, "Speed Up", uuid.New())
var TaijutsuUpSource = NewStatUpSource(game.StatTaijutsu, "Taijutsu Up", uuid.New())
