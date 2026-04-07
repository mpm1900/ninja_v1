import { MODIFIER_CLASSES, MODIFIER_ICONS } from '#/data/icons'
import type { Actor } from '#/lib/game/actor'
import type { Modifier } from '#/lib/game/modifier'
import { cn } from '#/lib/utils'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'

function ActorModifier({
  count,
  modifier,
}: {
  count: number
  modifier: Modifier | undefined
}) {
  if (!modifier || !modifier.name) return null
  const Icon = modifier?.icon ? MODIFIER_ICONS[modifier.icon] : undefined
  return (
    <span>
      {Icon ? (
        <Tooltip>
          <TooltipTrigger asChild>
            <Icon
              className={cn(
                'size-6 my-1',
                modifier?.icon && MODIFIER_CLASSES[modifier.icon]
              )}
            />
          </TooltipTrigger>
          <TooltipContent>{modifier?.name}</TooltipContent>
        </Tooltip>
      ) : (
        modifier?.name
      )}
      {count > 1 ? ` (${count})` : null}
    </span>
  )
}

function ActorModifiers({
  actor,
  modifiers,
}: {
  actor: Actor
  modifiers: Modifier[]
}) {
  return (
    <div className="relative flex flex-row-reverse justify-end flex-wrap px-2 gap-2 z-30">
      {Object.entries(actor.applied_modifiers ?? {}).map(([ID, count]) => (
        <ActorModifier
          key={ID}
          count={count}
          modifier={modifiers.find((m) => m.group_ID === ID)}
        />
      ))}
    </div>
  )
}

export { ActorModifiers }
