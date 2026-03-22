package game

import "github.com/google/uuid"

type Mutation[I any, O any, C any] struct {
	Delta    func(input I, context *C) O    `json:"-"`
	Filter   func(input I, context *C) bool `json:"-"`
	Priority int                            `json:"priority"`
}

type Transaction[M any, C any] struct {
	ID       uuid.UUID `json:"ID"`
	Context  *C        `json:"context"`
	Priority int       `json:"priority"`
	Mutation *M        `json:"mutation"`
}

func MakeTransaction[M any, C any](
	mutation *M,
	context *C,
) Transaction[M, C] {
	return Transaction[M, C]{
		ID:       uuid.New(),
		Context:  context,
		Mutation: mutation,
	}
}

func CheckTransaction[I any, O any, C any](
	input I,
	transaction *Transaction[Mutation[I, O, C], C],
) bool {
	if transaction == nil || transaction.Mutation == nil {
		return false
	}
	if transaction.Mutation.Filter == nil {
		return true
	}
	return transaction.Mutation.Filter(input, transaction.Context)
}

func ResolveTransaction[I any, O any, C any](
	input I,
	transaction *Transaction[Mutation[I, O, C], C],
	fallback O,
) (O, bool) {
	if !CheckTransaction(input, transaction) {
		return fallback, false
	}

	return transaction.Mutation.Delta(input, transaction.Context), true
}

func ResolveTransactionFn[I any, O any, C any](
	input I,
	transaction *Transaction[Mutation[I, O, C], C],
	fallback func(I) O,
) (O, bool) {
	return ResolveTransaction(input, transaction, fallback(input))
}
