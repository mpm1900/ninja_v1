import { Store } from '@tanstack/store'
import { NULL_CONTEXT, type Context } from '../game/context'
import type { Game } from '../game/game'

const battleContext = new Store<
  Context & { previous_action_IDs: Record<string, string> }
>({
  ...NULL_CONTEXT,
  previous_action_IDs: {},
})

function setContextPlayer(player_ID: string) {
  battleContext.setState((s) => ({
    ...s,
    source_player_ID: player_ID,
    parent_actor_ID: null,
    source_actor_ID: null,
    target_actor_IDs: [],
    target_position_IDs: [],
    previous_action_IDs: {},
  }))
}

function setActionID(actor_ID: string, action_ID: string) {
  battleContext.setState((s) => ({
    ...s,
    parent_actor_ID: null,
    source_actor_ID: null,
    target_actor_IDs: [],
    target_position_IDs: [],
    previous_action_IDs: {
      ...s.previous_action_IDs,
      [actor_ID]: action_ID,
    },
  }))
}

function setContextSource(source_ID: string, game: Game) {
  const existing = game.actions.find(
    (t) => t.context.source_actor_ID === source_ID
  )

  if (existing) {
    return battleContext.setState((s) => ({
      ...s,
      ...existing.context,
    }))
  }

  const source = game.actors.find((a) => a.ID === source_ID)
  if (!source) return

  const prev_ID = battleContext.get().previous_action_IDs[source_ID]
  const prev = source.actions.find((a) => a.ID === prev_ID)

  battleContext.setState((s) => ({
    ...s,
    action_ID: prev?.ID ?? source?.actions[0].ID ?? null,
    parent_actor_ID: source_ID,
    source_actor_ID: source_ID,
    target_actor_IDs: [],
    target_position_IDs: [],
  }))
}

function setContextAction(action_ID: string) {
  battleContext.setState((s) => ({
    ...s,
    action_ID,
    target_actor_IDs: [],
    target_position_IDs: [],
  }))
}

function setContext(context: Context) {
  battleContext.setState((s) => ({
    ...s,
    ...context,
  }))
}


export {
  battleContext,
  setContext,
  setContextPlayer,
  setActionID,
  setContextSource,
  setContextAction,
}
