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

type Transaction[M any] struct {
	ID       uuid.UUID `json:"ID"`
	Ready    bool      `json:"ready"`
	Context  Context   `json:"context"`
	Priority int       `json:"priority"`
	Mutation M         `json:"mutation"`
}

func MakeTransaction[M any](
	mutation M,
	context Context,
) Transaction[M] {
	return Transaction[M]{
		ID:       uuid.New(),
		Context:  context,
		Mutation: mutation,
		Ready:    true,
	}
}

func CheckTransaction[I any, O any](
	input I,
	transaction Transaction[Mutation[I, O]],
) bool {
	if transaction.Mutation.Filter == nil {
		return true
	}
	return transaction.Mutation.Filter(input, transaction.Context)
}

func ResolveTransaction[I any, O any](
	input I,
	transaction Transaction[Mutation[I, O]],
	fallback O,
) (O, bool) {
	if !CheckTransaction(input, transaction) {
		return fallback, false
	}

	return transaction.Mutation.Delta(input, transaction.Context), true
}

func ResolveTransactionFn[I any, O any](
	input I,
	transaction Transaction[Mutation[I, O]],
	fallback func(I) O,
) (O, bool) {
	return ResolveTransaction(input, transaction, fallback(input))
}

func NewGameMutation() GameMutation {
	return GameMutation{
		Delta:    func(input Game, context Context) Game { return input },
		Filter:   func(input Game, context Context) bool { return true },
		Priority: 0,
	}
}

func ComposeMutations(mutations ...GameMutation) GameMutation {
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
