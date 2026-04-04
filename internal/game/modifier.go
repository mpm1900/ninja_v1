package game

import (
	"github.com/google/uuid"
)

const ModifierDurationInf = -1

type Modifier struct {
	ID       uuid.UUID  `json:"ID"`
	GroupID  *uuid.UUID `json:"group_ID"`
	Name     string     `json:"name"`
	Duration int        `json:"duration"`

	Mutations []ModifierMutation `json:"-"`
	Triggers  []Trigger          `json:"triggers"`
}

func ResolveTrigger(game Game, transaction Transaction[Trigger]) []Transaction[GameMutation] {
	return transaction.Mutation.Delta(game, transaction.Context)
}

func CheckModifierForActor(tx Transaction[Modifier], game Game, actor Actor) bool {
	game.Actors = []Actor{actor}
	for _, mut := range tx.Mutation.Mutations {
		if mut.Filter(game, tx.Context) {
			return true
		}
	}

	return false
}
