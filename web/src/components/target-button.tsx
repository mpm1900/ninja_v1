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
}: {
  actor: Actor
  context: Context
  contextValid: boolean
  enabled: boolean
  loading: boolean
  onContextChange: (context: Context) => void
}) {
  const includes = context.target_actor_IDs.includes(actor.ID)
  return (
    <Button
      disabled={loading || (contextValid && !includes) || !enabled}
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

export { TargetButton }
