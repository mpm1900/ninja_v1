package instance

import (
	"context"
	"maps"
	"slices"
	"time"

	"ninja_v1/internal/game"

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
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		ReadRequest: make(chan Request),

		Game: game.NewGame(),
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

func (i *Instance) Run() {
	for {
		select {
		case client := <-i.Register:
			i.RegisterClient(client)
			i.BroadcastClients()

			i.Game.AddPlayer(game.NewPlayer(client.ID, 2))
			i.BroadcastGame()

			i.PostRegister(client)
		case client := <-i.Unregister:
			i.UnregisterClient(client)
			i.BroadcastClients()

			i.Game.RemovePlayer(client.ID)
			i.BroadcastGame()
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

func (i *Instance) RunGameActions() {
	go func() {
		i.Game.Status = game.GameStatusRunning
		time.Sleep(time.Second / 2)
		for i.Game.Next() {
			i.BroadcastGame()
			time.Sleep(time.Second)
		}

		i.Game.Status = game.GameStatusIdle
		i.BroadcastGame()
	}()
}
