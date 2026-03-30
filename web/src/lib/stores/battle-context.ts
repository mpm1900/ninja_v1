import { Store } from '@tanstack/store'
import { NULL_CONTEXT, type Context } from '../game/context'

const battleContext = new Store<Context>(NULL_CONTEXT)

function setContextPlayer(player_ID: string) {
  battleContext.setState((s) => ({
    ...s,
    source_player_ID: player_ID,
    parent_actor_ID: null,
    source_actor_ID: null,
    target_actor_IDs: [],
    target_position_IDs: [],
  }))
}

function setContextSource(source_ID: string, action_ID: string | null = null) {
  battleContext.setState((s) => ({
    ...s,
    action_ID,
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

export { battleContext, setContextPlayer, setContextSource, setContextAction }
