package game

import (
	"github.com/google/uuid"
)

type Trigger struct {
	Action[Game]
	On string `json:"on"`
}

type ModifierMutation struct {
	ActorMutation
	ModifierID *uuid.UUID `json:"modifier_ID"`
}

type Modifier struct {
	ID       uuid.UUID `json:"ID"`
	Name     string    `json:"name"`
	Duration *int      `json:"duration"`

	Mutations []ModifierMutation
	Triggers  []Trigger
}

type ModifierTransaction struct {
	ID       uuid.UUID `json:"ID"`
	Context  *Context  `json:"context"`
	Mutation Modifier  `json:"mutation"`
}

func MakeModifier(name string) Modifier {
	id := uuid.New()
	modifier := Modifier{
		ID:   id,
		Name: name,
	}

	return modifier
}
