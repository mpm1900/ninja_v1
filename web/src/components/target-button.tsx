import type { Action } from '#/lib/game/action'
import type { Actor } from '#/lib/game/actor'
import type { Context } from '#/lib/game/context'
import { Button } from './ui/button'

function TargetButton({
  actor,
  context,
  contextValid,
  enabled,
  loading,
  onContextChange,
  targetType,
}: {
  actor: Actor
  context: Context
  contextValid: boolean
  enabled: boolean
  loading: boolean
  onContextChange: (context: Context) => void
  targetType: Action['target_type']
}) {
  const includes =
    targetType === 'target-actor-id'
      ? context.target_actor_IDs?.includes(actor.ID)
      : context.target_position_IDs?.includes(actor.position_ID)

  return (
    <Button
      disabled={loading || (contextValid && !includes) || !enabled}
      variant={includes ? 'default' : 'ghost'}
      onClick={() => {
        if (targetType === 'target-actor-id') {
          onContextChange({
            ...context,
            target_actor_IDs: includes
              ? context.target_actor_IDs?.filter((id) => id !== actor.ID) ?? null
              : [...context.target_actor_IDs ?? [], actor.ID],
          })
        }

        if (targetType === 'target-position-type') {
          onContextChange({
            ...context,
            target_position_IDs: includes
              ? context.target_position_IDs?.filter(
                (id) => id !== actor.position_ID
              ) ?? null
              : [...(context.target_position_IDs ?? []), actor.position_ID],
          })
        }
      }}
    >
      {actor.name}
    </Button>
  )
}

export { TargetButton }
