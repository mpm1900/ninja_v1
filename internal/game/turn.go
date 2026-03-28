package game

type Turn struct {
	Count int
}

func NewTurn() Turn {
	return Turn{
		Count: 0,
	}
}
