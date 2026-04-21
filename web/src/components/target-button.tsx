import { SHINOBI_ICONS } from '#/data/icons'
import type { Action } from '#/lib/game/action'
import { type Actor } from '#/lib/game/actor'
import type { Context } from '#/lib/game/context'
import { addHoverTarget, removeHoverTarget } from '#/lib/stores/battle-context'
import { cn } from '#/lib/utils'
import { MiniHealthBar } from './actor-thumbnail'
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
      className="relative flex-col h-auto p-2 px-3 min-w-30 w-auto overflow-hidden shadow-[4px_4px_8px_rgba(0,0,0,0.5)]"
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
        console.log(targetType)
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
      <div className={cn("flex items-end w-full justify-between gap-4 relative z-10", !includes && "text-shadow-[1px_1px_0px_#000000]")}>
        <div>{actor.name}</div>
        <div className='text-xs'>Lv<span className='font-black'>{actor.level}</span></div>
      </div>

      <div className="relative w-full">
        <MiniHealthBar actor={actor} className="left-0 right-0" />
      </div>
      <div className='absolute z-0 opacity-30 -left-4 -top-3'>
        {actor.affiliations?.filter((_, i) => i == 0).map((a) => {
          const C = SHINOBI_ICONS[a]
          return C ? <C key={a} className="w-12" /> : null
        })}
      </div>
    </Button>
  )
}

export { TargetButton }
