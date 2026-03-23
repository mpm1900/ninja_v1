package instance

import (
	"context"
	"maps"
	"slices"

	"ninja_v1/internal/game"
	data "ninja_v1/internal/game/data"

	"github.com/google/uuid"
)

type Instance struct {
	ID      uuid.UUID             `json:"ID"`
	ctx     context.Context       `json:"-"`
	Clients map[uuid.UUID]*Client `json:"clients,omitempty"`
	Game    game.Game             `json:"game"`

	Register    chan *Client `json:"-"`
	Unregister  chan *Client `json:"-"`
	ReadRequest chan Request `json:"-"`
}

func NewInstance(ctx context.Context, id uuid.UUID) *Instance {
	return &Instance{
		ID:          id,
		ctx:         ctx,
		Clients:     make(map[uuid.UUID]*Client),
		Game:        game.Game{},
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		ReadRequest: make(chan Request),
	}
}

func (i *Instance) RegisterClient(client *Client) {
	i.Clients[client.ID] = client
}

func (i *Instance) UnregisterClient(client *Client) {
	delete(i.Clients, client.ID)
}

func (i *Instance) BroadcastGame() {
	// fmt.Printf("BROADCAST STATE %#v\n", game)
	for _, client := range i.Clients {
		select {
		case client.inbox <- NewStateMessage(&i.Game):
		// if a client is unable to handle the state update,
		//   unregister them so they don't the loop
		default:
			i.UnregisterClient(client)
		}
	}
}

func (i *Instance) PostRegister(client *Client) {
	client.inbox <- PostRegisterMssage(client, &i.Game)
}

func (i *Instance) BroadcastClients() {
	clients := slices.Collect(maps.Values(i.Clients))
	for _, client := range i.Clients {
		select {
		case client.inbox <- NewClientsMessage(clients):
		// if a client is unable to handle the state update,
		//   unregister them so they don't the loop
		default:
			i.UnregisterClient(client)
		}
	}
}

const (
	state = iota
	clients
	none
)

func Reducer(instance *Instance, request Request) int {
	switch request.Type {
	case "add-actor":
		def := data.ACTORS[*request.Context.SourceActorID]
		actor := game.MakeActor(def, request.ClientID, 13824)
		instance.Game.AddActor(actor)
		return state
	case "remove-actor":
		index := slices.IndexFunc(instance.Game.Actors, func(a game.Actor) bool {
			return a.ID == *request.Context.SourceActorID
		})

		if index == -1 {
			return none
		}

		instance.Game.RemoveActor(*request.Context.SourceActorID)
		return state

	case "add-modifier":
		if request.ModifierID == nil {
			return none
		}

		if modifier, ok := data.MODIFIERS[*request.ModifierID]; ok {
			transaction := game.MakeModifierTransaction(&modifier, &request.Context)
			instance.Game.AddModifier(transaction)
			return state
		}

		return none
	case "remove-modifier":
		instance.Game.FilterModifiers(func(m game.Transaction[game.Modifier]) bool {
			return m.ID != *request.ModifierID
		})
		return state

	default:
		return none
	}
}

func (i *Instance) Run() {
	for {
		select {
		case client := <-i.Register:
			i.RegisterClient(client)
			i.BroadcastClients()
			i.PostRegister(client)
		case client := <-i.Unregister:
			i.UnregisterClient(client)
			i.BroadcastClients()
		case request := <-i.ReadRequest:
			switch Reducer(i, request) {
			case state:
				i.BroadcastGame()
			case clients:
				i.BroadcastClients()
			case none:
			}
		}
	}
}
