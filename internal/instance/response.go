package instance

import (
	"ninja_v1/internal/game"
)

type Response struct {
	Type    string     `json:"type"`
	State   *game.GameJSON `json:"state"`
	Clients []*Client  `json:"clients"`
}

func NewStateMessage(client *Client, state *game.Game) Response {
	var json game.GameJSON
	if (client == nil) {
		json = state.ToJSON(nil)
	} else {
		json = state.ToJSON(&client.ID)
	}

	return Response{
		Type:    "state",
		State:   &json,
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
	var json game.GameJSON
	if (client == nil) {
		json = state.ToJSON(nil)
	} else {
		json = state.ToJSON(&client.ID)
	}

	return Response{
		Type:    "join-success",
		State:   &json,
		Clients: []*Client{client},
	}
}
