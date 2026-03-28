package game

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/google/uuid"
)

type GameMutation = Mutation[Game, Game]
type GameTransaction = Transaction[GameMutation]

type GameStatus string
type GameLog = string

const (
	GameStatusRunning GameStatus = "running"
	GameStatusIdle    GameStatus = "idle"
	GameStatusWaiting GameStatus = "waiting"
)

/**
 * Game is the main state container state for a game instance
 */
type Game struct {
	Status    GameStatus              `json:"status"`
	Turn      Turn                    `json:"turn"`
	Players   []Player                `json:"players"`
	Actors    []Actor                 `json:"actors"`
	Modifiers []Transaction[Modifier] `json:"modifiers"`

	/*
	 * [Transactions] are the quable/storable change in state
	 * 	- mostly to animate and modify state changes while running
	 *  - so one aciton can do multiple things, in order
	 */
	Transactions Queue[GameTransaction] `json:"transactions"`
	/*
	 * [Triggers] are basically Actions that do not need input to resolve
	 *  - triggers are to delay, repeat, or statically/conditionally respond to events
	 */
	Triggers Queue[Transaction[Trigger]] `json:"triggers"`
	/*
	 * [Actions] are a sorted list of actions to be resolved
	 *  - actions resolve down to reansactions
	 *
	 * [Prompts] are special actions that pause Action resolution loops.
	 *  - can stop actions from resolving while waiting for user input
	 *  - switch ins will be a common Prompt
	 */
	Actions Queue[Transaction[Action]] `json:"actions"`
	Prompts Queue[Transaction[Action]]

	Log []GameLog `json:"log"`
}

func NewGame() Game {
	return Game{
		Status:       GameStatusIdle,
		Players:      make([]Player, 0),
		Actors:       make([]Actor, 0),
		Modifiers:    make([]Transaction[Modifier], 0),
		Transactions: MakeQueue[GameTransaction](),
		Actions:      MakeQueue[Transaction[Action]](),
		Prompts:      MakeQueue[Transaction[Action]](),
		Triggers:     MakeQueue[Transaction[Trigger]](),
		Log:          []string{},
	}
}

/**
 * Getters
 */

// Player
func (g Game) GetPlayer(predicate func(Player) bool) (bool, Player) {
	for _, a := range g.Players {
		if predicate(a) {
			return true, a
		}
	}

	return false, Player{}
}
func (g Game) GetPlayerByID(ID uuid.UUID) (bool, Player) {
	return g.GetPlayer(func(p Player) bool {
		return p.ID == ID
	})
}

// Actor
func (g Game) GetActor(predicate func(Actor) bool) (bool, Actor) {
	for _, a := range g.Actors {
		if predicate(a) {
			return true, a
		}
	}

	return false, Actor{}
}
func (g Game) GetActorByID(ID uuid.UUID) (bool, Actor) {
	return g.GetActor(func(a Actor) bool {
		return a.ID == ID
	})
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
func (g Game) GetActorsByPlayer(playerID uuid.UUID) []Actor {
	return g.GetActors(func(a Actor) bool {
		return a.PlayerID == playerID
	})
}
func (g Game) GetActionableActors() []Actor {
	return g.GetActors(func(a Actor) bool {
		active := a.PositionID != nil
		return active && a.Alive && !a.Stunned
	})
}
func (g Game) GetActorsFilters(filters ...ActorFilter) []Actor {
	filter := ComposeAF(filters...)
	return g.GetActors(func(a Actor) bool {
		return filter(a, Context{})
	})
}

func (g Game) GetResolvedActors() map[uuid.UUID]ResolvedActor {
	actors := make(map[uuid.UUID]ResolvedActor, len(g.Actors))
	for _, a := range g.Actors {
		resolvedActor := a.Resolve(g)
		actors[a.ID] = resolvedActor
	}

	return actors
}

func (g Game) GetTriggers(context Context) []Transaction[Trigger] {
	triggers := []Transaction[Trigger]{
		MakeTransaction(END_OF_TURN_TRIGGER, context),
	}
	modifiers := make([]Transaction[Modifier], 0, len(g.Modifiers))
	modifiers = append(modifiers, g.Modifiers...)
	modifiers = append(modifiers, GetActorModifiers(g)...)

	for _, mod := range modifiers {
		for _, trig := range mod.Mutation.Triggers {
			if trig.Check != nil && !trig.Check(g, context, mod) {
				continue
			}
			triggers = append(triggers, MakeTransaction(trig, context))
		}
	}

	return triggers
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

func (g *Game) FilterParentModifiers(actorID uuid.UUID) {
	g.FilterModifiers(func(modifier Transaction[Modifier]) bool {
		if modifier.Context.ParentActorID != nil {
			return *modifier.Context.ParentActorID != actorID
		}

		return true
	})
}

func (g *Game) AddPlayer(player Player) {
	g.Players = append(g.Players, player)
}

func (g *Game) RemovePlayer(playerID uuid.UUID) {
	g.Players = slices.DeleteFunc(g.Players, func(p Player) bool {
		return p.ID == playerID
	})

	for i, _ := range g.Actors {
		g.Actors[i].PositionID = nil
	}
}

func (g *Game) UpdatePlayer(playerID uuid.UUID, updater func(Player) Player) {
	index := slices.IndexFunc(g.Players, func(p Player) bool {
		return p.ID == playerID
	})

	if index == -1 {
		return
	}

	g.Players[index] = updater(g.Players[index])
}

func (g *Game) SetPosition(actor Actor, positionID *uuid.UUID) {
	prev := actor.PositionID
	ok, player := g.GetPlayerByID(actor.PlayerID)
	if !ok {
		return
	}

	if positionID == nil {
		g.FilterParentModifiers(actor.ID)
	}

	if positionID != nil {
		curr := player.GetActorAtPosition(*positionID)
		if curr != nil {
			ok, displaced := g.GetActorByID(*curr)
			if ok {
				g.SetPosition(displaced, nil)
			}
		}

		g.UpdatePlayer(actor.PlayerID, func(p Player) Player {
			p.SetPosition(*positionID, &actor.ID)
			return p
		})
	}

	g.UpdateActor(actor.ID, func(a Actor) Actor {
		a.PositionID = positionID
		return a
	})

	if prev != nil {
		g.UpdatePlayer(actor.PlayerID, func(p Player) Player {
			p.SetPosition(*prev, nil)
			return p
		})
	}
}

func (g *Game) SetActorPlayerIndex(actor Actor, index *int) {
	if index == nil {
		g.SetPosition(actor, nil)
	}

	ok, player := g.GetPlayerByID(actor.PlayerID)
	if !ok || *index >= len(player.Positions) {
		return
	}

	positionID := player.Positions[*index]
	g.SetPosition(actor, &positionID.ID)
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

func (g *Game) SetActionCooldown(actorID uuid.UUID, actionID uuid.UUID, cooldown int) {
	g.UpdateActor(actorID, func(a Actor) Actor {
		a.SetActionCooldown(actionID, cooldown)
		return a
	})
}

func (g *Game) PushAction(transaction Transaction[Action]) bool {
	for _, t := range g.Actions {
		if *t.Context.SourceActorID == *transaction.Context.SourceActorID {
			return false
		}
	}

	g.Actions.Enqueue(transaction)
	slices.SortFunc(g.Actions, func(a, b Transaction[Action]) int {
		if a.Mutation.Priority != b.Mutation.Priority {
			return b.Mutation.Priority - a.Mutation.Priority
		}

		ok, a_source := g.GetActorByID(*a.Context.SourceActorID)
		if !ok {
			return 1
		}
		ok, b_source := g.GetActorByID(*b.Context.SourceActorID)
		if !ok {
			return -1
		}

		a_res := a_source.Resolve(*g)
		b_res := b_source.Resolve(*g)
		return b_res.Stats[StatSpeed] - a_res.Stats[StatSpeed]
	})

	if len(g.Actions) == len(g.GetActionableActors()) {
		return true
	}

	return false
}

func (g Game) HasPlayerPrompt(playerID uuid.UUID) bool {
	for _, prompt := range g.Prompts {
		if *prompt.Context.SourcePlayerID == playerID {
			return true
		}
	}

	return false
}

func (g Game) AllPromptsReady() bool {
	if len(g.Prompts) == 0 {
		return false
	}

	for _, prompt := range g.Prompts {
		if !prompt.Ready {
			return false
		}
	}

	return true
}

func (g *Game) AddPrompt(transaction Transaction[Action]) {
	g.Prompts = append(g.Prompts, transaction)
}

func (g *Game) RemovePrompt(ID uuid.UUID) {
	g.Prompts = slices.DeleteFunc(g.Prompts, func(t Transaction[Action]) bool {
		return t.ID == ID
	})
}

func (g *Game) ReadyPrompt(ID uuid.UUID, context Context) {
	for i, prompt := range g.Prompts {
		if prompt.ID == ID {
			g.Prompts[i].Context = context
			g.Prompts[i].Ready = true
			return
		}
	}
}

func (g *Game) RunPrompt(transaction Transaction[Action]) {
	transactions := ResolveAction(g, transaction)
	g.Transactions = append(g.Transactions, transactions...)
}

func (g *Game) RunAction(transaction Transaction[Action]) {
	if transaction.Context.SourceActorID != nil {
		_, source := g.GetActorByID(*transaction.Context.SourceActorID)
		g.PushLog(fmt.Sprintf("%s used %s.", source.Name, transaction.Mutation.Config.Name))

		cost := transaction.Mutation.Cost
		if cost.Delta != nil {
			costTx := MakeTransaction(cost, transaction.Context)
			g.Transactions = append(g.Transactions, costTx)
		}

	}

	transactions := ResolveAction(g, transaction)
	g.Transactions = append(g.Transactions, transactions...)
}

func (g *Game) RunTrigger(transaction Transaction[Trigger]) {
	transactions := ResolveTrigger(*g, transaction)
	g.Transactions = append(g.Transactions, transactions...)
}

func (g *Game) On(on TriggerOn, context Context) {
	triggers := make([]Transaction[Trigger], 0)
	for _, trigger := range g.GetTriggers(context) {
		if trigger.Mutation.On == on {
			triggers = append(triggers, trigger)
		}
	}

	g.Triggers = append(g.Triggers, triggers...)
}

func (g *Game) JumpTransaction(transaction Transaction[GameMutation]) {
	next := Queue[GameTransaction]{transaction}
	g.Transactions = append(next, g.Transactions...)
}

func (g *Game) PushLog(log GameLog) {
	g.Log = append(g.Log, log)
}

func (g *Game) NextTurn() {
	g.Turn.Count++
}

type GameJSON struct {
	Status  GameStatus      `json:"status"`
	Players []Player        `json:"players"`
	Actors  []ResolvedActor `json:"actors"`

	Modifiers    []Transaction[Modifier]     `json:"modifiers"`
	Transactions []Transaction[GameMutation] `json:"transactions"`
	Actions      []Transaction[Action]       `json:"actions"`
	Prompt       *Transaction[Action]        `json:"prompt"`
	Triggers     []Transaction[Trigger]      `json:"triggers"`

	Log []GameLog `json:"log"`
}

func (g Game) ToJSON(playerID *uuid.UUID) GameJSON {
	resolvedMap := g.GetResolvedActors()
	resolved := make([]ResolvedActor, 0, len(g.Actors))
	for _, a := range g.Actors {
		resolved = append(resolved, resolvedMap[a.ID])
	}

	var prompt *Transaction[Action]
	if playerID != nil {
		for _, p := range g.Prompts {
			if *p.Context.SourcePlayerID == *playerID && p.Ready == false {
				prompt = &p
				break
			}
		}
	}

	status := g.Status
	if status == GameStatusIdle && len(g.Prompts) > 0 && prompt == nil {
		status = GameStatusWaiting
	}

	return GameJSON{
		Status:       status,
		Players:      g.Players,
		Actors:       resolved,
		Modifiers:    g.Modifiers,
		Transactions: g.Transactions,
		Actions:      g.Actions,
		Prompt:       prompt,
		Triggers:     g.Triggers,
		Log:          g.Log,
	}
}

func (g Game) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.ToJSON(nil))
}
