import type { Action } from "#/lib/game/action"
import type { Actor } from "#/lib/game/actor"
import type { Context } from "#/lib/game/context"
import { useEffect, useState } from "react"

function useGameContext(actor: Actor | undefined, action_ID?: string, deps: unknown[] = []) {
  const [context, onContextChange] = useState<Context>({
    action_ID: action_ID ?? null,
    source_player_ID: actor?.player_ID ?? null,
    source_actor_ID: actor?.ID ?? null,
    parent_actor_ID: actor?.ID ?? null,
    target_actor_IDs: [],
    target_position_IDs: [],
  })

  useEffect(() => {
    onContextChange(c => ({
      ...c,
      source_actor_ID: actor?.ID ?? null,
      parent_actor_ID: actor?.ID ?? null,
    }))
  }, [actor?.ID])

  useEffect(() => {
    onContextChange(c => ({
      ...c,
      action_ID: action_ID ?? null,
      target_actor_IDs: [],
      target_position_IDs: [],
    }))
  }, [action_ID])

  useEffect(() => {
    onContextChange(c => ({
      ...c,
      target_actor_IDs: [],
      target_position_IDs: [],
    }))
  }, deps)

  return {
    context,
    onContextChange,
  }
}

export { useGameContext }
