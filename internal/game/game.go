package game

import (
	"encoding/json"
	"slices"
)

type GameTransaction struct {
	Transaction[Game, Game, Context]
}

type Game struct {
	Actors    []Actor               `json:"actors"`
	Modifiers []ModifierTransaction `json:"modifiers"`

	Transactions []GameTransaction         `json:"transactions"`
	Actions      []ActionTransaction[Game] `json:"actions"`
	Trigger      []ActionTransaction[Game] `json:"triggers"`
}

func (g *Game) SetActors(actors []Actor) {
	g.Actors = slices.Clone(actors)
}

func (g *Game) AddModifier(modifier ModifierTransaction) {
	g.Modifiers = append(g.Modifiers, modifier)
}

func (g *Game) FilterModifiers(predicate func(modifier ModifierTransaction) bool) {
	filtered := g.Modifiers[:0]
	for _, m := range g.Modifiers {
		if predicate(m) {
			filtered = append(filtered, m)
		}
	}

	g.Modifiers = filtered
}

func (g Game) MarshalJSON() ([]byte, error) {
	actorModifiers := GetActorModifiers(g)
	resolved := make([]ResolvedActor, 0, len(g.Actors))

	for _, a := range g.Actors {
		resolvedActor := ResolveActor(a, g.Modifiers, actorModifiers)
		resolved = append(resolved, resolvedActor)
	}

	type gameJSON struct {
		Actors    []ResolvedActor       `json:"actors"`
		Modifiers []ModifierTransaction `json:"modifiers"`

		Transactions []GameTransaction         `json:"transactions"`
		Actions      []ActionTransaction[Game] `json:"actions"`
		Trigger      []ActionTransaction[Game] `json:"triggers"`
	}

	return json.Marshal(gameJSON{
		Actors:       resolved,
		Modifiers:    g.Modifiers,
		Transactions: g.Transactions,
		Actions:      g.Actions,
		Trigger:      g.Trigger,
	})
}
