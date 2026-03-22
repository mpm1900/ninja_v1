package game

import (
	"github.com/google/uuid"
)

type Trigger struct {
	Action
	On string `json:"on"`
}

type ModifierMutation struct {
	ActorMutation
	ModifierID    *uuid.UUID
	TransactionID *uuid.UUID
}

type Modifier struct {
	ID       uuid.UUID `json:"ID"`
	Name     string    `json:"name"`
	Duration *int      `json:"duration"`

	Mutations []ModifierMutation `json:"-"`
	Triggers  []Trigger          `json:"triggers"`
}

func MakeModifier(name string) Modifier {
	id := uuid.New()
	modifier := Modifier{
		ID:   id,
		Name: name,
	}

	return modifier
}

func MakeModifierTransaction(modifier *Modifier, context *Context) Transaction[Modifier, Context] {
	return MakeTransaction(modifier, context)

}
