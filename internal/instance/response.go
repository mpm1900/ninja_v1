package instance

import (
	"ninja_v1/internal/game"
)

type Response struct {
	Type    string     `json:"type"`
	State   *game.Game `json:"state"`
	Clients []*Client  `json:"clients"`
}

func NewStateMessage(state *game.Game) Response {
	return Response{
		Type:    "state",
		State:   state,
		Clients: nil,
	}
}

func NewClientsMessage(clients []*Client) Response {
	return Response{
		Type:    "clients",
		State:   nil,
		Clients: clients,
	}
}

func PostRegisterMssage(client *Client, state *game.Game) Response {
	return Response{
		Type:    "join-success",
		State:   state,
		Clients: []*Client{client},
	}
}
