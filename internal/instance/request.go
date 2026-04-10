package instance

import (
	"ninja_v1/internal/game"
	"ninja_v1/internal/game/data"

	"github.com/google/uuid"
)

type RequestType = string

const (
	SetTeam        RequestType = "set-team"
	AddActor       RequestType = "add-actor"
	RemoveActor    RequestType = "remove-actor" //TEMP
	UpdateActor    RequestType = "update-actor"
	PushAction     RequestType = "push-action"
	RemoveAction   RequestType = "remove-action"
	RunGameActions RequestType = "run-game-actions" // TEMP
	ValidateState  RequestType = "validate-state"   // TEMP
	ResolvePrompt  RequestType = "resolve-prompt"

	GetTargets      RequestType = "get-targets"
	ValidateContext RequestType = "validate-context"
)

type ActorConfig struct {
	AbilityID *uuid.UUID             `json:"ability_ID"`
	ActionIDs []uuid.UUID            `json:"action_IDs"`
	Focus     *game.ActorFocus       `json:"focus"`
	ItemID    *uuid.UUID             `json:"item_ID"`
	AuxStats  map[game.ActorStat]int `json:"aux_stats"`
}

type TeamActor struct {
	ActorID uuid.UUID   `json:"actor_ID"`
	Config  ActorConfig `json:"config"`
}

type TeamConfig struct {
	Actors []TeamActor `json:"actors"`
	Name   string      `json:"name"`
}

type Request struct {
	Type        RequestType  `json:"type"`
	ClientID    uuid.UUID    `json:"client_ID"`
	PromptID    *uuid.UUID   `json:"prompt_ID"`
	Context     game.Context `json:"context"`
	ActorConfig *ActorConfig `json:"actor_config"`
	TeamConfig  *TeamConfig  `json:"team_config"`
}

type HydratedActorConfig struct {
	Ability  *game.Modifier
	Actions  []game.Action
	Focus    game.ActorFocus
	Item     *game.Modifier
	AuxStats map[game.ActorStat]int `json:"aux_stats"`
}

func HydrateActorConfig(config ActorConfig, abilities []game.Modifier) HydratedActorConfig {
	var ability *game.Modifier = nil
	if config.AbilityID != nil {
		for _, a := range abilities {
			if a.ID == *config.AbilityID {
				ability = &a
			}
		}
	} else {
		if len(abilities) > 0 {
			ability = &abilities[0]
		}
	}

	actions := make([]game.Action, 0, len(config.ActionIDs))
	for _, id := range config.ActionIDs {
		action, ok := data.ACTIONS[id]
		if !ok {
			continue
		}
		actions = append(actions, action)
	}

	var item *game.Modifier = nil
	if config.ItemID != nil {
		i, ok := data.ITEMS[*config.ItemID]
		if ok {
			item = &i
		}
	}

	focus := game.FocusNone
	if config.Focus != nil {
		focus = *config.Focus
	}

	return HydratedActorConfig{
		Ability:  ability,
		Actions:  actions,
		AuxStats: config.AuxStats,
		Focus:    focus,
		Item:     item,
	}
}
