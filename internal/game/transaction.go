package game

import "github.com/google/uuid"

type Mutation[I any, O any] struct {
	Delta    func(input I, context *Context) O    `json:"-"`
	Filter   func(input I, context *Context) bool `json:"-"`
	Priority int                                  `json:"priority"`
}

type Transaction[M any] struct {
	ID       uuid.UUID `json:"ID"`
	Context  *Context  `json:"context"`
	Priority int       `json:"priority"`
	Mutation *M        `json:"mutation"`
}

func MakeTransaction[M any](
	mutation *M,
	context *Context,
) Transaction[M] {
	return Transaction[M]{
		ID:       uuid.New(),
		Context:  context,
		Mutation: mutation,
	}
}

func CheckTransaction[I any, O any](
	input I,
	transaction *Transaction[Mutation[I, O]],
) bool {
	if transaction == nil || transaction.Mutation == nil {
		return false
	}
	if transaction.Mutation.Filter == nil {
		return true
	}
	return transaction.Mutation.Filter(input, transaction.Context)
}

func ResolveTransaction[I any, O any](
	input I,
	transaction *Transaction[Mutation[I, O]],
	fallback O,
) (O, bool) {
	if !CheckTransaction(input, transaction) {
		return fallback, false
	}

	return transaction.Mutation.Delta(input, transaction.Context), true
}

func ResolveTransactionFn[I any, O any](
	input I,
	transaction *Transaction[Mutation[I, O]],
	fallback func(I) O,
) (O, bool) {
	return ResolveTransaction(input, transaction, fallback(input))
}
