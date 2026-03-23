package game

import (
	"encoding/json"
	"slices"

	"github.com/google/uuid"
)

type GameMutation = Mutation[Game, Game]

type Game struct {
	Actors    []Actor                 `json:"actors"`
	Modifiers []Transaction[Modifier] `json:"modifiers"`

	Transactions []Transaction[GameMutation] `json:"transactions"`
	Actions      []Transaction[Action]       `json:"actions"`
	Trigger      []Transaction[Action]       `json:"triggers"`
}

func (g Game) GetActor(predicate func(Actor) bool) (bool, Actor) {
	for _, a := range g.Actors {
		if predicate(a) {
			return true, a
		}
	}

	return false, Actor{}
}

func (g Game) GetActors(predicate func(Actor) bool) []Actor {
	actors := make([]Actor, 0)
	for _, a := range g.Actors {
		if predicate(a) {
			actors = append(actors, a)
		}
	}
	return actors
}

func (g *Game) FilterModifiers(predicate func(modifier Transaction[Modifier]) bool) {
	filtered := g.Modifiers[:0]
	for _, m := range g.Modifiers {
		if predicate(m) {
			filtered = append(filtered, m)
		}
	}

	g.Modifiers = filtered
}

func (g *Game) AddModifier(modifier Transaction[Modifier]) {
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

func (g *Game) UpdateActor(actorID uuid.UUID, updater func(Actor) Actor) {
	index := slices.IndexFunc(g.Actors, func(a Actor) bool {
		return a.ID == actorID
	})

	if index == -1 {
		return
	}

	g.Actors[index] = updater(g.Actors[index])
}

func (g Game) MarshalJSON() ([]byte, error) {
	resolved := make([]ResolvedActor, 0, len(g.Actors))

	for _, a := range g.Actors {
		resolvedActor := ResolveActor(a, g)
		resolved = append(resolved, resolvedActor)
	}

	type gameJSON struct {
		Actors []ResolvedActor `json:"actors"`

		Modifiers    []Transaction[Modifier]     `json:"modifiers"`
		Transactions []Transaction[GameMutation] `json:"transactions"`
		Actions      []Transaction[Action]       `json:"actions"`
		Trigger      []Transaction[Action]       `json:"triggers"`
	}

	return json.Marshal(gameJSON{
		Actors:       resolved,
		Modifiers:    g.Modifiers,
		Transactions: g.Transactions,
		Actions:      g.Actions,
		Trigger:      g.Trigger,
	})
}
