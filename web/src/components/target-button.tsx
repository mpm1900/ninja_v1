import type { Action } from '#/lib/game/action'
import { getVitals, type Actor } from '#/lib/game/actor'
import type { Context } from '#/lib/game/context'
import { addHoverTarget, removeHoverTarget } from '#/lib/stores/battle-context'
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
  const vitals = getVitals(actor)
  const includes =
    targetType === 'target-actor-id'
      ? context.target_actor_IDs?.includes(actor.ID)
      : context.target_position_IDs?.includes(actor.position_ID)

  return (
    <Button
      disabled={loading || (contextValid && !includes) || !enabled}
      variant={
        includes
          ? 'default'
          : context.source_player_ID === actor.player_ID
            ? 'player_target'
            : 'enemy_target'
      }
      onMouseEnter={() => {
        addHoverTarget(actor.ID)
      }}
      onMouseLeave={() => {
        removeHoverTarget(actor.ID)
      }}
      onClick={() => {
        if (targetType === 'target-actor-id') {
          onContextChange({
            ...context,
            target_actor_IDs: includes
              ? (context.target_actor_IDs?.filter((id) => id !== actor.ID) ??
                null)
              : [...(context.target_actor_IDs ?? []), actor.ID],
          })
        }

        if (targetType === 'target-position-type') {
          onContextChange({
            ...context,
            target_position_IDs: includes
              ? (context.target_position_IDs?.filter(
                  (id) => id !== actor.position_ID
                ) ?? null)
              : [...(context.target_position_IDs ?? []), actor.position_ID],
          })
        }
      }}
    >
      {actor.name} ({(vitals.hp.ratio * 100).toFixed()}%)
    </Button>
  )
}

export { TargetButton }
