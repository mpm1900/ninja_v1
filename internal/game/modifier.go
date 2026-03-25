package game

import (
	"github.com/google/uuid"
)

type TriggerOn string

const (
	OnDamageRecieve TriggerOn = "on-damage-recieve"
)

var TRIGGERS []TriggerOn = []TriggerOn{
	OnDamageRecieve,
}

type Trigger struct {
	ActionMutation
	ID    uuid.UUID                                       `json:"ID"`
	On    TriggerOn                                       `json:"on"`
	Check func(Game, Context, Transaction[Modifier]) bool `json:"-"`
}

type ModifierMutation struct {
	Mutation[Actor, Actor]
	ModifierGroupID *uuid.UUID
	TransactionID   *uuid.UUID
}

type Modifier struct {
	ID       uuid.UUID `json:"ID"`
	GroupID  uuid.UUID `json:"group_ID"`
	Name     string    `json:"name"`
	Duration *int      `json:"duration"`

	Mutations []ModifierMutation `json:"-"`
	Triggers  []Trigger          `json:"triggers"`
}

func ResolveTrigger(game Game, transaction Transaction[Trigger]) []Transaction[GameMutation] {
	return transaction.Mutation.Delta(game, transaction.Context)
}

func MakeModifierMutation(
	modifierGroupID *uuid.UUID,
	priority int,
	filter func(input Actor, context Context) bool,
	delta func(input Actor, context Context) Actor,
) ModifierMutation {
	return ModifierMutation{
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
