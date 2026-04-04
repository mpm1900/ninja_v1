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

func AddLogs(logs ...GameLog) GameMutation {
	return GameMutation{
		Delta: func(g Game, context Context) Game {
			g.Log = append(g.Log, logs...)

			return g
		},
	}
}

/**
 * Modifier Mutations
 * [GameMutation]
 */
type ModifierMutation struct {
	GameMutation
	ActorFilter     func(Game, Actor, Context) bool
	ActorDelta      func(Game, Actor, Context) Actor
	ModifierGroupID *uuid.UUID
	TransactionID   *uuid.UUID
}

func MakeActorMutation(
	modifierGroupID *uuid.UUID,
	priority int,
	filter func(Game, Actor, Context) bool,
	delta func(Game, Actor, Context) Actor,
) ModifierMutation {
	return ModifierMutation{
		ActorFilter:     filter,
		ActorDelta:      delta,
		ModifierGroupID: modifierGroupID,
		GameMutation: GameMutation{
			Filter: func(g Game, context Context) bool {
				for _, actor := range g.Actors {
					if filter(g, actor, context) {
						return true
					}
				}
				return false
			},
			Delta: func(g Game, context Context) Game {
				for _, actor := range g.Actors {
					if filter(g, actor, context) {
						g.UpdateActor(actor.ID, func(a Actor) Actor {
							return delta(g, a, context)
						})
					}
				}
				return g
			},
			Priority: priority,
		},
	}
}
