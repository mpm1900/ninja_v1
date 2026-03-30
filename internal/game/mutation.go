package game

import (
	"slices"

	"github.com/google/uuid"
)

type Mutation[I any, O any] struct {
	Delta    func(input I, context Context) O    `json:"-"`
	Filter   func(input I, context Context) bool `json:"-"`
	Priority int                                 `json:"priority"`
}

/**
 * Game Mutations
 * [GameMutation]
 */
type GameMutation = Mutation[Game, Game]

func NewGameMutation() GameMutation {
	return GameMutation{
		Delta:    func(input Game, context Context) Game { return input },
		Filter:   func(input Game, context Context) bool { return true },
		Priority: 0,
	}
}

func ComposeGameMutations(mutations ...GameMutation) GameMutation {
	if len(mutations) == 0 {
		return NewGameMutation()
	}

	slices.SortFunc(mutations, func(a, b GameMutation) int {
		return b.Priority - a.Priority
	})

	return GameMutation{
		Delta: func(input Game, context Context) Game {
			for _, mut := range mutations {
				if mut.Delta != nil {
					input = mut.Delta(input, context)
				}
			}

			return input
		},
		Filter: func(input Game, context Context) bool {
			for _, mut := range mutations {
				if mut.Filter != nil && !mut.Filter(input, context) {
					return false
				}
			}
			return true
		},
		Priority: mutations[0].Priority,
	}
}

/**
 * Actor Mutations
 * [ActorMutation]
 */

type ActorMutation struct {
	Mutation[Actor, Actor]
	ModifierGroupID *uuid.UUID
	TransactionID   *uuid.UUID
}

func MakeActorMutation(
	modifierGroupID *uuid.UUID,
	priority int,
	filter func(input Actor, context Context) bool,
	delta func(input Actor, context Context) Actor,
) ActorMutation {
	return ActorMutation{
		ModifierGroupID: modifierGroupID,
		Mutation: Mutation[Actor, Actor]{
			Filter:   filter,
			Delta:    delta,
			Priority: priority,
		},
	}
}
