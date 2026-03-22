package game

import (
	"github.com/google/uuid"
)

type Trigger struct {
	Action
	On string `json:"on"`
}

type ModifierMutation struct {
	Mutation[Actor, Actor, Context]
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

func MakeModifierMutation(
	modifierID *uuid.UUID,
	priority int,
	filter func(input Actor, context *Context) bool,
	delta func(input Actor, context *Context) Actor,
) ModifierMutation {
	return ModifierMutation{
		ModifierID: modifierID,
		Mutation: Mutation[Actor, Actor, Context]{
			Filter:   filter,
			Delta:    delta,
			Priority: priority,
		},
	}
}

func MakeModifierTransaction(modifier *Modifier, context *Context) Transaction[Modifier, Context] {
	return MakeTransaction(modifier, context)
}
