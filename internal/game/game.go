package game

type GameTransaction struct {
	Transaction[Game, Game, Context]
}

type Game struct {
	Actors    []Actor `json:"actors"`
	Modifiers []ModifierTransaction

	Transactions []GameTransaction
	Actions      []ActionTransaction[Game]
	Trigger      []ActionTransaction[Game]
}
