package game

import "github.com/google/uuid"

type TriggerOn string

const (
	OnDamageRecieve TriggerOn = "on-damage-recieve"
	OnTurnEnd TriggerOn = "on-turn-end"
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
	ID: uuid.New(),
	On: OnTurnEnd,
	Check: func(g Game, context Context, tx Transaction[Modifier]) bool {return true},
	ActionMutation: ActionMutation{
		Delta: func(input Game, context Context) []Transaction[GameMutation] {
			var transactions []Transaction[GameMutation]
			mut := GameMutation{
				Delta: func(g Game, c Context) Game {
					for i, _ := range g.Actors {
						g.Actors[i].DecrementCooldowns()
					}

					for i, mod := range g.Modifiers {
						if mod.Mutation.Duration != nil {
							*g.Modifiers[i].Mutation.Duration -= 1
						}
					}

					g.FilterModifiers(func(mod Transaction[Modifier]) bool {
						if mod.Mutation.Duration != nil {
							return *mod.Mutation.Duration <= 0
						}
						return true
					})

					return g
				},
			}

			tx := MakeTransaction(mut, context)
			transactions = append(transactions, tx)

			return transactions
		},
	},
}
