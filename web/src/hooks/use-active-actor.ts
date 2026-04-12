import { battleContext } from "#/lib/stores/battle-context"
import { gameStore } from "#/lib/stores/game"
import { useStore } from "@tanstack/react-store"

function useActiveActor() {
  const context = useStore(battleContext, (c) => c)
  const actors = useStore(gameStore, g => g.actors)
  const actor = actors.find((a) => a.ID === context.source_actor_ID)
  return actor
}

export { useActiveActor }
