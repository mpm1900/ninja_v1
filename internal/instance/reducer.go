package instance

import (
	"fmt"
	"ninja_v1/internal/game"
	data "ninja_v1/internal/game/data"
	"slices"
)

func Reducer(instance *Instance, request Request) int {
	switch request.Type {
	case AddActor:
		def, ok := data.ACTORS[*request.Context.SourceActorID]
		if !ok {
			fmt.Println("[AddActor] Unknown Actor")
			return none
		}

		ok, player := instance.Game.GetPlayerByID(request.ClientID)
		if !ok {
			fmt.Println("[AddActor] Unknown Player")
			return none
		}

		actors := instance.Game.GetActorsByPlayer(player.ID)
		if len(actors) >= player.TeamCapacity {
			fmt.Println("[AddActor] Team Full")
			return none
		}
		actor := game.MakeActor(def, request.ClientID, 13824, data.ACTIONS)
		instance.Game.AddActor(actor)
		instance.Game.PushLog(fmt.Sprintf("Actor joined: %s", actor.Name))
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

	case AddModifier:
		if request.ModifierID == nil {
			return none
		}

		if modifier, ok := data.MODIFIERS[*request.ModifierID]; ok {
			transaction := game.MakeModifierTransaction(modifier, request.Context)
			instance.Game.AddModifier(transaction)
			return state
		}

		return none
	case RemoveModifier:
		instance.Game.FilterModifiers(func(m game.Transaction[game.Modifier]) bool {
			return m.ID != *request.ModifierID
		})
		return state

	case PushAction:
		if request.Context.ActionID == nil {
			fmt.Println("no context action_ID")
			return none
		}

		action, ok := data.ACTIONS[*request.Context.ActionID]
		if !ok {
			fmt.Println("action not found")
			return none
		}

		transaction := game.MakeTransaction(action, request.Context)
		if instance.Game.PushAction(transaction) {
			instance.RunGameActions()
		}

		return state
	case RemoveAction:
		if request.Context.ActionID == nil {
			fmt.Println("no context action_ID")
			return none
		}

		instance.Game.Actions = slices.DeleteFunc(instance.Game.Actions, func(tx game.Transaction[game.Action]) bool {
			return tx.ID == *request.Context.ActionID
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

	case SetActorPlayer:
		targets := instance.Game.GetTargets(request.Context)

		if len(targets) == 0 {
			return none
		}

		target := targets[0]
		ok, player := instance.Game.GetPlayerByID(*request.Context.SourcePlayerID)
		if !ok {
			return none
		}

		instance.Game.SetPosition(target, nil)
		instance.Game.UpdatePlayer(player.ID, func(p game.Player) game.Player {
			p.AddPosition(&target.ID)
			return p
		})
		instance.Game.UpdateActor(target.ID, func(a game.Actor) game.Actor {
			a.PlayerID = player.ID
			return a
		})

		return state

	case SetActorPosition:
		targets := instance.Game.GetTargets(request.Context)

		if len(targets) == 0 || request.PositionIndex == nil {
			return none
		}

		target := targets[0]
		instance.Game.SetActorPlayerIndex(target, request.PositionIndex)
		return state

	default:
		return none
	}

}
