import { MODIFIER_CLASSES, MODIFIER_ICONS } from '#/data/icons'
import type { Actor } from '#/lib/game/actor'
import type { Modifier } from '#/lib/game/modifier'
import { cn } from '#/lib/utils'
import { ModifierTooltip } from './modifier-tooltip'

function ActorModifier({
  count,
  modifier,
}: {
  count: number
  modifier: Modifier | undefined
}) {
  if (!modifier || !modifier.show) return null
  const Icon = modifier?.icon ? MODIFIER_ICONS[modifier.icon] : undefined
  return (
    <ModifierTooltip modifier={modifier} contentProps={{ sideOffset: 0 }}>
      <span className="text-shadow-[1px_1px_0px_#000000] cursor-default">
        {Icon ? (
          <Icon
            className={cn(
              'size-6 my-1',
              modifier?.icon && MODIFIER_CLASSES[modifier.icon]
            )}
          />
        ) : (
          modifier?.name
        )}
        {count > 1 ? ` (${count})` : null}
      </span>
    </ModifierTooltip>
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
    <div className="relative flex flex-row-reverse justify-end items-end flex-wrap px-2 gap-2 z-30">
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
