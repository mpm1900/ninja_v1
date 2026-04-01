package game

import (
	"github.com/google/uuid"
)

type TriggerOn string

const (
	OnActorEnter    TriggerOn = "on-actor-enter"
	OnActorLeave    TriggerOn = "on-actor-leave"
	OnDamageRecieve TriggerOn = "on-damage-recieve"
	OnTurnEnd       TriggerOn = "on-turn-end"
)

var TRIGGERS []TriggerOn = []TriggerOn{
	OnDamageRecieve,
	OnTurnEnd,
}

type Trigger struct {
	ActionMutation
	ID    uuid.UUID                                       `json:"ID"`
	On    TriggerOn                                       `json:"on"`
	Check func(Game, Context, Transaction[Modifier]) bool `json:"-"`
}

var END_OF_TURN_TRIGGER Trigger = Trigger{
	ID:    uuid.New(),
	On:    OnTurnEnd,
	Check: func(g Game, context Context, tx Transaction[Modifier]) bool { return true },
	ActionMutation: ActionMutation{
		Delta: func(input Game, context Context) []Transaction[GameMutation] {
			var transactions []Transaction[GameMutation]
			mut := GameMutation{
				Delta: func(g Game, c Context) Game {
					t := g.Turn.Count
					for i, _ := range g.Actors {
						if t > 0 {
							g.Actors[i].DecrementCooldowns()
							g.Actors[i].RecoverStamina(g, 0.08)
						}
						g.Actors[i].ActiveTurns++
					}

					if t > 0 {
						g.FilterModifiers(func(mod Transaction[Modifier]) bool {
							return mod.Mutation.Duration != 0
						})

						for i, _ := range g.Modifiers {
							g.Modifiers[i].Mutation.Duration -= 1
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
