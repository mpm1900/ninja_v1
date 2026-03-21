package modifiers

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

func NewStatDoublerSource(stat game.BaseStat, name string) game.Modifier {
	statDoublerID := uuid.New()
	modifier := game.Modifier{
		ID:   statDoublerID,
		Name: name,
		Mutations: []game.ModifierMutation{
			{
				ModifierID: &statDoublerID,
				ActorMutation: game.ActorMutation{
					Filter: func(actor game.Actor, context *game.Context) bool {
						return actor.ID == *context.SourceActorID
					},
					Delta: func(actor game.Actor, context *game.Context) game.Actor {
						actor.Stages[stat] = actor.Stages[stat] + 2
						return actor
					},
				},
			},
		}}

	return modifier
}
