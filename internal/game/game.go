package game

type GameTransaction struct {
	Transaction[Game, Game, Context]
}

type Game struct {
	Actors    []Actor               `json:"actors"`
	Modifiers []ModifierTransaction `json:"modifiers"`

	Transactions []GameTransaction         `json:"transactions"`
	Actions      []ActionTransaction[Game] `json:"actions"`
	Trigger      []ActionTransaction[Game] `json:"triggers"`
}
