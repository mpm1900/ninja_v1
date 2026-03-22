package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewStatDoublerSource(stat game.BaseStat, name string) game.Modifier {
	statDoublerID := uuid.New()
	mut := game.MakeModifierMutation(
		&statDoublerID,
		game.PriorityDefault,
		game.SourceFilter,
		func(actor game.Actor, context *game.Context) game.Actor {
			actor.Stages[stat] = actor.Stages[stat] + 2
			return actor
		},
	)

	return game.Modifier{
		ID:   statDoublerID,
		Name: name,
		Mutations: []game.ModifierMutation{
			mut,
		},
	}
}
