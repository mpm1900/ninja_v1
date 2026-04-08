package instance

import (
	"context"
	"maps"
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data"
	"slices"
	"time"

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

		Game: game.NewGame(data.ACTIONS),
	}
}

func (i *Instance) RegisterClient(client *Client) {
	if existing, ok := i.Clients[client.ID]; ok && existing != client {
		existing.cancel()
	}
	i.Clients[client.ID] = client
}

func (i *Instance) UnregisterClient(client *Client) bool {
	existing, ok := i.Clients[client.ID]
	if !ok || existing != client {
		return false
	}

	delete(i.Clients, client.ID)
	return true
}

func (i *Instance) BroadcastGame() {
	// fmt.Printf("BROADCAST STATE %#v\n", game)
	for _, client := range i.Clients {
		select {
		case client.inbox <- NewGameMessage(client, &i.Game):
		// if a client is unable to handle the state update,
		//   unregister them so they don't the loop
		default:
			i.UnregisterClient(client)
		}
	}
}

func (i *Instance) PostRegister(client *Client) {
	client.inbox <- PostRegisterMessage(client, &i.Game)
}

func (i *Instance) TargetIDsResponse(clientID uuid.UUID, context game.Context, targetIDs []uuid.UUID) {
	client, ok := i.Clients[clientID]
	if !ok {
		return
	}

	client.inbox <- TargetIDsResponse(client, context, targetIDs)
}
func (i *Instance) ValidateContextResponse(clientID uuid.UUID, context game.Context, valid bool) {
	client, ok := i.Clients[clientID]
	if !ok {
		return
	}

	client.inbox <- ValidateContextMessage(client, context, valid)
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

			if _, ok := i.Game.GetPlayerByID(client.ID); !ok {
				i.Game.AddPlayer(game.NewPlayer(client.ID, 2, *client.User))
				i.BroadcastGame()
			}

			i.PostRegister(client)
		case client := <-i.Unregister:
			removed := i.UnregisterClient(client)
			if !removed {
				continue
			}

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

func (i *Instance) RunGameActions() {
	go func() {
		i.Game.Status = game.GameStatusRunning
		i.BroadcastGame()

	resolveStep:
		for {
			for i.Game.Next() {
				i.BroadcastGame()
				time.Sleep(i.Game.Tick)
			}

			if len(i.Game.Prompts) > 0 {
				break resolveStep
			}

			switch i.Game.Turn.Phase {
			case game.TurnMain:
				// Main is complete for this run, proceed to End.
				i.Game.NextPhase()
				i.BroadcastGame()
				continue

			case game.TurnEnd:
				if i.Game.Turn.Count > 0 {
					i.Game.On(game.OnTurnEnd, nil)
				}

				for i.Game.Next() {
					i.BroadcastGame()
					time.Sleep(i.Game.Tick)
				}

				if len(i.Game.Prompts) > 0 {
					break resolveStep
				}

				// End is complete, proceed to Cleanup.
				i.Game.NextPhase()
				i.BroadcastGame()
				continue

			case game.TurnCleanup:
				time.Sleep(i.Game.Tick)
				// Cleanup is complete, advance turn and reset to Main.
				i.Game.NextTurn()
				i.BroadcastGame()
				break resolveStep

			default:
				// Recover unknown/setup phases by moving toward Main.
				i.Game.NextPhase()
				i.BroadcastGame()
			}
		}

		i.Game.Status = game.GameStatusIdle
		i.BroadcastGame()
	}()
}
