package game

import (
	"encoding/json"
	"maps"
	"slices"

	"github.com/google/uuid"
)

type GameMutation = Mutation[Game, Game]
type GameTransaction = Transaction[GameMutation]

type GameStatus string

type Player = uuid.UUID

const (
	GameStatusRunning GameStatus = "running"
	GameStatusIdle    GameStatus = "idle"
)

type Game struct {
	Status    GameStatus              `json:"status"`
	Players   map[Player]struct{}     `json:"players"`
	Actors    []Actor                 `json:"actors"`
	Modifiers []Transaction[Modifier] `json:"modifiers"`

	Transactions Queue[GameTransaction]     `json:"transactions"`
	Actions      Queue[Transaction[Action]] `json:"actions"`
	Triggers     Queue[Transaction[Action]] `json:"triggers"`
}

func NewGame() Game {
	return Game{
		Status:       GameStatusIdle,
		Players:      make(map[Player]struct{}),
		Actors:       make([]Actor, 0),
		Modifiers:    make([]Transaction[Modifier], 0),
		Transactions: MakeQueue[GameTransaction](),
		Actions:      MakeQueue[Transaction[Action]](),
		Triggers:     MakeQueue[Transaction[Action]](),
	}
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

func (g *Game) AddPlayer(player Player) {
	g.Players[player] = struct{}{}
}

func (g *Game) RemovePlayer(player Player) {
	delete(g.Players, player)
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

func (g *Game) ExecuteTransaction(transaction Transaction[Action]) {
	transactions := ResolveAction(*g, transaction)
	g.Transactions = append(g.Transactions, transactions...)
}

func (g *Game) Flush() {
	for g.Next() {
	}
}

func (g *Game) Next() bool {
	transaction, err := g.Transactions.Dequeue()
	if err != nil {
		return false
	}

	n, ok := ResolveTransaction(*g, transaction, *g)
	if !ok {
		return false
	}

	*g = n
	return true
}

func (g Game) MarshalJSON() ([]byte, error) {
	resolved := make([]ResolvedActor, 0, len(g.Actors))

	for _, a := range g.Actors {
		resolvedActor := ResolveActor(a, g)
		resolved = append(resolved, resolvedActor)
	}

	type gameJSON struct {
		Status  GameStatus      `json:"status"`
		Players []Player        `json:"players"`
		Actors  []ResolvedActor `json:"actors"`

		Modifiers    []Transaction[Modifier]     `json:"modifiers"`
		Transactions []Transaction[GameMutation] `json:"transactions"`
		Actions      []Transaction[Action]       `json:"actions"`
		Triggers     []Transaction[Action]       `json:"triggers"`
	}

	return json.Marshal(gameJSON{
		Status:       g.Status,
		Players:      slices.Collect(maps.Keys(g.Players)),
		Actors:       resolved,
		Modifiers:    g.Modifiers,
		Transactions: g.Transactions,
		Actions:      g.Actions,
		Triggers:     g.Triggers,
	})
}
