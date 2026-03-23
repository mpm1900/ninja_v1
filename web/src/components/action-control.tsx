import type { Context } from '#/lib/game/context'
import { useQuery } from '@tanstack/react-query'
import { Button } from './ui/button'
import { useStore } from '@tanstack/react-store'
import { socketStore } from '#/lib/stores/socket'
import { gameStore } from '#/lib/stores/game'
import { actionTargetsQuery } from '#/lib/queries/action-targets'
import { isActionContextValidQuery } from '#/lib/queries/is-action-context-valid'
import { useEffect } from 'react'
import type { Actor } from '#/lib/game/actor'

function TargetButton({
  actor,
  context,
  contextValid,
  loading,
  onContextChange,
}: {
  actor: Actor
  context: Context
  contextValid: boolean
  loading: boolean
  onContextChange: (context: Context) => void
}) {
  const includes = context.target_actor_IDs.includes(actor.ID)
  return (
    <Button
      disabled={loading || (contextValid && !includes)}
      variant={includes ? 'default' : 'ghost'}
      onClick={() => {
        onContextChange({
          ...context,
          target_actor_IDs: includes
            ? context.target_actor_IDs.filter((id) => id !== actor.ID)
            : [...context.target_actor_IDs, actor.ID],
        })
      }}
    >
      {actor.name}
    </Button>
  )
}

function ActionControl({
  action_ID,
  context,
  onContextChange,
}: {
  action_ID: string
  context: Context
  onContextChange: (context: Context) => void
}) {
  const instanceID = useStore(socketStore, (s) => s.instanceID!)
  const game = useStore(gameStore, (g) => g)
  const valid = useQuery(isActionContextValidQuery(action_ID, context))
  const actionTargets = useQuery(
    actionTargetsQuery(instanceID, action_ID, context)
  )
  const loading = valid.isFetching || actionTargets.isFetching

  useEffect(() => {
    actionTargets.refetch()
    onContextChange({
      ...context,
      target_actor_IDs: [],
    })
  }, [game.actors.length])

  return (
    <div>
      <div className="p-2 flex gap-2">
        {game.actors
          .filter((a) => actionTargets.data?.includes(a.ID))
          .map((a) => (
            <TargetButton
              key={a.ID}
              actor={a}
              loading={loading}
              contextValid={!!valid.data}
              context={context}
              onContextChange={onContextChange}
            />
          ))}
      </div>
      <Button disabled={loading || !valid.data}>Select</Button>
    </div>
  )
}

export { ActionControl }
