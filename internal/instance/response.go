package instance

import (
	"ninja_v1/internal/game"

	"github.com/google/uuid"
)

type ResponseType string

const (
	ResponseTypeGame            = "game"
	ResponseTypeClients         = "clients"
	ResponseTypeJoinSuccess     = "join-success"
	ResponseTypeValidateContext = "validate-context"
	ResponseTypeTargetIDs       = "target-IDs"
)

type Response struct {
	Type      ResponseType   `json:"type"`
	State     *game.GameJSON `json:"state"`
	Clients   []*Client      `json:"clients"`
	Valid     *bool          `json:"valid"`
	Context   *game.Context  `json:"context"`
	TargetIDs []uuid.UUID    `json:"target_IDs"`
}

func NewGameMessage(client *Client, state *game.Game) Response {
	var json game.GameJSON
	if client == nil {
		json = state.ToJSON(nil)
	} else {
		json = state.ToJSON(&client.ID)
	}

	return Response{
		Type:    ResponseTypeGame,
		State:   &json,
		Clients: nil,
	}
}

func NewClientsMessage(clients []*Client) Response {
	return Response{
		Type:    ResponseTypeClients,
		State:   nil,
		Clients: clients,
	}
}

func PostRegisterMessage(client *Client, state *game.Game) Response {
	var json game.GameJSON
	if client == nil {
		json = state.ToJSON(nil)
	} else {
		json = state.ToJSON(&client.ID)
	}

	return Response{
		Type:    ResponseTypeJoinSuccess,
		State:   &json,
		Clients: []*Client{client},
	}
}

func TargetIDsResponse(client *Client, context game.Context, targetIDs []uuid.UUID) Response {
	return Response{
		Type:      ResponseTypeTargetIDs,
		Context:   &context,
		TargetIDs: targetIDs,
	}
}

func ValidateContextMessage(client *Client, context game.Context, valid bool) Response {
	return Response{
		Type:    ResponseTypeValidateContext,
		Context: &context,
		Valid:   &valid,
	}
}
