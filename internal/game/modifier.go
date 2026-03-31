package game

import (
	"github.com/google/uuid"
)

const ModifierDurationInf = -1

type Modifier struct {
	ID       uuid.UUID `json:"ID"`
	GroupID  uuid.UUID `json:"group_ID"`
	Name     string    `json:"name"`
	Duration int       `json:"duration"`

	Mutations []ActorMutation `json:"-"`
	Triggers  []Trigger       `json:"triggers"`
}

func ResolveTrigger(game Game, transaction Transaction[Trigger]) []Transaction[GameMutation] {
	return transaction.Mutation.Delta(game, transaction.Context)
}

func MakeModifierTransaction(modifier Modifier, context Context) Transaction[Modifier] {
	return MakeTransaction(modifier, context)
}
