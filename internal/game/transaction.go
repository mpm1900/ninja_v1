package game

type Mutation[I any, O any, C any] struct {
	Delta    func(input I, context *C) O    `json:"-"`
	Filter   func(input I, context *C) bool `json:"-"`
	Priority int                            `json:"priority"`
}

type Transaction[I any, O any, C any] struct {
	ID       string             `json:"ID"`
	Context  *C                 `json:"context"`
	Priority int                `json:"priority"`
	Mutation *Mutation[I, O, C] `json:"-"`
}

func MakeTransaction[I any, O any, C any](
	mutation *Mutation[I, O, C],
	context *C) Transaction[I, O, C] {
	return Transaction[I, O, C]{
		ID:       "",
		Context:  context,
		Mutation: mutation,
	}
}

func CheckTransaction[I any, O any, C any](
	input I,
	transaction *Transaction[I, O, C],
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
	transaction *Transaction[I, O, C],
	fallback O,
) (O, bool) {
	if !CheckTransaction(input, transaction) {
		return fallback, false
	}

	return transaction.Mutation.Delta(input, transaction.Context), true
}

func ResolveTransactionFn[I any, O any, C any](
	input I,
	transaction *Transaction[I, O, C],
	fallback func(I) O,
) (O, bool) {
	return ResolveTransaction(input, transaction, fallback(input))
}
