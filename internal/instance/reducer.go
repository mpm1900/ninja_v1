package instance

import (
	"ninja_v1/internal/game"
	data "ninja_v1/internal/game/data"
	"slices"
	"time"
)

func Reducer(instance *Instance, request Request) int {
	switch request.Type {
	case AddActor:
		def, ok := data.ACTORS[*request.Context.SourceActorID]
		if !ok {
			return none
		}

		actor := game.MakeActor(def, request.ClientID, 13824)
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
		action, ok := data.ACTIONS[*request.ActionID]
		if !ok {
			return none
		}

		instance.Game.Status = game.GameStatusRunning
		transaction := game.MakeTransaction(action, request.Context)
		instance.Game.Actions.Enqueue(transaction)

		go func() {
			time.Sleep(time.Second / 2)
			for instance.Game.Next() {
				instance.BroadcastGame()
				time.Sleep(time.Second)
			}

			instance.Game.Status = game.GameStatusIdle
			instance.BroadcastGame()
		}()

		return state

	case SetActorPlayer:
		count, targets := instance.Game.GetTargets(request.Context)

		if count == 0 {
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
		count, targets := instance.Game.GetTargets(request.Context)

		if count == 0 || request.PositionIndex == nil {
			return none
		}

		target := targets[0]
		instance.Game.SetActorPlayerIndex(target, request.PositionIndex)
		return state

	default:
		return none
	}

}
