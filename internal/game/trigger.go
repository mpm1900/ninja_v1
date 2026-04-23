package game

import (
	"github.com/google/uuid"
)

type TriggerOn string

const (
	OnActorEnter    TriggerOn = "on-actor-enter"
	OnActorLeave    TriggerOn = "on-actor-leave"
	OnActionStart   TriggerOn = "on-action-start"
	OnActionEnd     TriggerOn = "on-action-end"
	OnImmortalSave  TriggerOn = "on-immortal-save"
	OnDamageReceive TriggerOn = "on-damage-receive"
	OnDeath         TriggerOn = "on-death"
	OnKill          TriggerOn = "on-kill"
	OnModifierAdd   TriggerOn = "on-modifier-add"
	OnTurnEnd       TriggerOn = "on-turn-end"
)

var TRIGGERS []TriggerOn = []TriggerOn{
	OnDamageReceive,
	OnTurnEnd,
}

type Trigger struct {
	ActionMutation
	ID         uuid.UUID                                             `json:"ID"`
	ModifierID uuid.UUID                                             `json:"modifier_ID"`
	On         TriggerOn                                             `json:"on"`
	Check      func(Game, Game, Context, Transaction[Modifier]) bool `json:"-"`
}
