package game

import (
	"github.com/google/uuid"
)

type TriggerOn string

const (
	OnActorEnter    TriggerOn = "on-actor-enter"
	OnActorLeave    TriggerOn = "on-actor-leave"
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

var END_OF_TURN_TRIGGER Trigger = Trigger{
	ID:    uuid.MustParse("f63aefeb-02cf-4dbd-93f9-8f1908f99d4f"),
	On:    OnTurnEnd,
	Check: func(p, g Game, context Context, tx Transaction[Modifier]) bool { return true },
	ActionMutation: ActionMutation{
		Delta: func(parent Game, input Game, context Context) []Transaction[GameMutation] {
			var transactions []Transaction[GameMutation]
			mut := GameMutation{
				Delta: func(p Game, g Game, c Context) Game {
					t := g.Turn.Count
					for i := range g.Actors {
						if t > 0 {
							g.Actors[i].DecrementCooldowns()
							g.Actors[i].RecoverStamina(g, 0.08)
						}
						g.Actors[i].IncrementTurns()
					}

					if t > 0 {
						g.FilterModifiers(func(mod Transaction[Modifier]) bool {
							return mod.Mutation.Duration != 0
						})

						for i, _ := range g.Modifiers {
							g.Modifiers[i].Mutation.DecrementTimers()
						}
					}

					return g
				},
			}

			tx := MakeTransaction(mut, context)
			transactions = append(transactions, tx)

			return transactions
		},
	},
}
