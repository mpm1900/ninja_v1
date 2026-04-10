package instance

import (
	"maps"
	"ninja_v1/internal/game"
	data "ninja_v1/internal/game/data"
	"slices"

	"github.com/google/uuid"
)

func Reducer(instance *Instance, request Request) int {
	switch request.Type {
	case GetTargets:
		if request.Context.ActionID == nil && request.PromptID == nil {
			instance.TargetIDsResponse(request.ClientID, request.Context, nil)
			return none
		}

		var action = game.Action{}
		if request.PromptID != nil {
			tx, ok := instance.Game.GetPromptTxByID(*request.PromptID)
			if !ok {
				instance.TargetIDsResponse(request.ClientID, request.Context, nil)
				return none
			}
			action = tx.Mutation
		} else {
			actor, ok := instance.Game.GetSource(request.Context)
			if !ok {
				instance.ValidateContextResponse(request.ClientID, request.Context, false)
				return none
			}

			a, ok := actor.GetActionByID(instance.Game, *request.Context.ActionID)
			if !ok {
				// if it's not on the source, do a root look up
				// this is probably due to something like recharing or a
				// non-stored, special action
				a, ok = data.ACTIONS[*request.Context.ActionID]
				if !ok {
					instance.TargetIDsResponse(request.ClientID, request.Context, nil)
					return none
				}
			}

			action = a
		}

		targets := instance.Game.GetActors(func(a game.Actor) bool {
			return action.TargetPredicate(instance.Game, a, request.Context)
		})
		targetIDs := make([]uuid.UUID, 0, len(targets))
		for _, a := range targets {
			targetIDs = append(targetIDs, a.ID)
		}

		instance.TargetIDsResponse(request.ClientID, request.Context, targetIDs)
		return none
	case ValidateContext:
		if request.Context.ActionID == nil && request.PromptID == nil {
			instance.ValidateContextResponse(request.ClientID, request.Context, false)
			return none
		}

		if request.PromptID != nil {
			action, ok := instance.Game.GetPromptTxByID(*request.PromptID)
			if !ok {
				instance.ValidateContextResponse(request.ClientID, request.Context, false)
				return none
			}

			valid := action.Mutation.ContextValidate(request.Context)
			instance.ValidateContextResponse(request.ClientID, request.Context, valid)
			return none
		}

		actor, ok := instance.Game.GetSource(request.Context)
		if !ok {
			instance.ValidateContextResponse(request.ClientID, request.Context, false)
			return none
		}

		action, ok := actor.GetActionByID(instance.Game, *request.Context.ActionID)
		if !ok {
			instance.ValidateContextResponse(request.ClientID, request.Context, false)
			return none
		}

		valid := action.ContextValidate(request.Context)
		instance.ValidateContextResponse(request.ClientID, request.Context, valid)
		return none
	case SetTeam:
		if request.TeamConfig == nil {
			return none
		}
		player, ok := instance.Game.GetPlayerByID(request.ClientID)
		if !ok {
			return none
		}

		if len(request.TeamConfig.Actors) > player.TeamCapacity {
			return none
		}

		teamActors := make([]game.Actor, len(request.TeamConfig.Actors))
		for i, actorConfig := range request.TeamConfig.Actors {
			def, ok := data.ACTORS[actorConfig.ActorID]
			if !ok {
				return none
			}

			hydrated := HydrateActorConfig(actorConfig.Config, def.Abilities)

			actor := game.MakeActor(
				def,
				request.ClientID,
				/* 24 13824 */ 1000000,
				hydrated.Ability,
				hydrated.Item,
				hydrated.Actions,
				hydrated.Focus,
				hydrated.AuxStats,
			)
			teamActors[i] = actor
		}
		instance.Game.SetPlayerActors(request.ClientID, teamActors)
		return state
	case AddActor:
		if request.Context.SourceActorID == nil {
			return none
		}

		def, ok := data.ACTORS[*request.Context.SourceActorID]
		if !ok {
			return none
		}

		config := ActorConfig{
			AbilityID: nil,
			ActionIDs: def.ActionIDs,
			AuxStats:  map[game.ActorStat]int{},
			ItemID:    nil,
			Focus:     nil,
		}

		hydrated := HydrateActorConfig(config, def.Abilities)
		player, ok := instance.Game.GetPlayerByID(request.ClientID)
		if !ok {
			return none
		}

		actors := instance.Game.GetActorsByPlayer(player.ID)
		if len(actors) >= player.TeamCapacity {
			return none
		}

		var ability *game.Modifier = nil
		if len(def.Abilities) > 0 {
			ability = &def.Abilities[0]
		}

		actor := game.MakeActor(
			def,
			request.ClientID,
			/* 24 13824 */ 1000000,
			ability,
			hydrated.Item,
			hydrated.Actions,
			hydrated.Focus,
			hydrated.AuxStats,
		)
		instance.Game.AddActor(actor)
		return state
	case RemoveActor:
		index := slices.IndexFunc(instance.Game.Actors, func(a game.Actor) bool {
			return a.ID == *request.Context.SourceActorID
		})

		if index == -1 {
			return none
		}

		actor := instance.Game.Actors[index]
		instance.Game.SetPosition(actor, nil)
		instance.Game.RemoveActor(actor.ID)
		return state
	case UpdateActor:
		if request.Context.SourceActorID == nil || request.ActorConfig == nil {
			return none
		}

		config := *request.ActorConfig
		instance.Game.UpdateActor(*request.Context.SourceActorID, func(a game.Actor) game.Actor {
			hydrated := HydrateActorConfig(config, a.Abilities)

			if config.ActionIDs != nil {
				a.SetActions(config.ActionIDs, data.ACTIONS)
			}
			a.Focus = hydrated.Focus
			if hydrated.Ability != nil {
				a.Ability = hydrated.Ability
			}
			if hydrated.Item != nil {
				a.Item = hydrated.Item
			}
			if hydrated.AuxStats != nil {
				a.AuxStats = maps.Clone(hydrated.AuxStats)
			}

			return a
		})

		return state

	case PushAction:
		if request.Context.ActionID == nil || request.Context.SourceActorID == nil {
			return none
		}

		actor, ok := instance.Game.GetSource(request.Context)
		if !ok {
			return none
		}

		action, ok := actor.GetActionByID(instance.Game, *request.Context.ActionID)
		if !ok {
			return none
		}

		if action.ID == game.Switch.ID && actor.SwitchLocked {
			return none
		}

		if action.ID != game.Switch.ID && actor.ActionLocked && actor.LastUsedActionID != nil {
			if *request.Context.ActionID != *actor.LastUsedActionID {
				return none
			}
		}

		transaction := game.MakeTransaction(action, request.Context)
		if instance.Game.PushAction(transaction) {
			instance.RunGameActions()
		}

		return state
	case RemoveAction: //CancelAction
		if request.Context.ActionID == nil {
			return none
		}

		instance.Game.Actions = slices.DeleteFunc(instance.Game.Actions, func(tx game.Transaction[game.Action]) bool {
			if tx.Context.SourcePlayerID != nil && *tx.Context.SourcePlayerID == request.ClientID {
				return tx.ID == *request.Context.ActionID
			}

			return false
		})

		return state

	case ResolvePrompt:
		if request.PromptID == nil {
			return none
		}

		instance.Game.ReadyPrompt(*request.PromptID, request.Context)
		if instance.Game.AllPromptsReady() {
			instance.RunGameActions()
		}

		return state

	case RunGameActions:
		if instance.Game.Status == game.GameStatusRunning {
			return none
		}

		instance.RunGameActions()
		return state
	case ValidateState:
		if instance.Game.Status == game.GameStatusRunning {
			return none
		}

		instance.RunGameActions()
		return state

	default:
		return none
	}

}
