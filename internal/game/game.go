package game

import (
	"encoding/json"
	"slices"

	"github.com/google/uuid"
)

type GameMutation struct {
	Mutation[Game, Game, Context]
}

type Game struct {
	Actors    []Actor                          `json:"actors"`
	Modifiers []Transaction[Modifier, Context] `json:"modifiers"`

	Transactions []Transaction[GameMutation, Context] `json:"transactions"`
	Actions      []Transaction[Action, Context]       `json:"actions"`
	Trigger      []Transaction[Action, Context]       `json:"triggers"`
}

func (g *Game) FilterModifiers(predicate func(modifier Transaction[Modifier, Context]) bool) {
	filtered := g.Modifiers[:0]
	for _, m := range g.Modifiers {
		if predicate(m) {
			filtered = append(filtered, m)
		}
	}

	g.Modifiers = filtered
}

func (g *Game) AddModifier(modifier Transaction[Modifier, Context]) {
	g.Modifiers = append(g.Modifiers, modifier)
}

func (g *Game) AddActor(actor Actor) {
	g.Actors = append(g.Actors, actor)
}

func (g *Game) RemoveActor(actorID uuid.UUID) {
	g.Actors = slices.DeleteFunc(g.Actors, func(a Actor) bool {
		return a.ID == actorID
	})
}

func (g Game) MarshalJSON() ([]byte, error) {
	resolved := make([]ResolvedActor, 0, len(g.Actors))

	for _, a := range g.Actors {
		resolvedActor := ResolveActor(a, g)
		resolved = append(resolved, resolvedActor)
	}

	type gameJSON struct {
		Actors []ResolvedActor `json:"actors"`

		Modifiers    []Transaction[Modifier, Context]     `json:"modifiers"`
		Transactions []Transaction[GameMutation, Context] `json:"transactions"`
		Actions      []Transaction[Action, Context]       `json:"actions"`
		Trigger      []Transaction[Action, Context]       `json:"triggers"`
	}

	return json.Marshal(gameJSON{
		Actors:       resolved,
		Modifiers:    g.Modifiers,
		Transactions: g.Transactions,
		Actions:      g.Actions,
		Trigger:      g.Trigger,
	})
}
