package game

import (
	"github.com/google/uuid"
)

type ActorMutation struct {
	Mutation[Actor, Actor]
	ModifierGroupID *uuid.UUID
	TransactionID   *uuid.UUID
}

type Modifier struct {
	ID       uuid.UUID `json:"ID"`
	GroupID  uuid.UUID `json:"group_ID"`
	Name     string    `json:"name"`
	Duration *int      `json:"duration"`

	Mutations []ActorMutation `json:"-"`
	Triggers  []Trigger       `json:"triggers"`
}

func ResolveTrigger(game Game, transaction Transaction[Trigger]) []Transaction[GameMutation] {
	return transaction.Mutation.Delta(game, transaction.Context)
}

func MakeActorMutation(
	modifierGroupID *uuid.UUID,
	priority int,
	filter func(input Actor, context Context) bool,
	delta func(input Actor, context Context) Actor,
) ActorMutation {
	return ActorMutation{
		ModifierGroupID: modifierGroupID,
		Mutation: Mutation[Actor, Actor]{
			Filter:   filter,
			Delta:    delta,
			Priority: priority,
		},
	}
}

func MakeModifierTransaction(modifier Modifier, context Context) Transaction[Modifier] {
	return MakeTransaction(modifier, context)
}
