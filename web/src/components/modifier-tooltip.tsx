import type React from 'react'
import { HoverCard, HoverCardContent, HoverCardTrigger } from './ui/hover-card'
import type { Modifier } from '#/lib/game/modifier'

function ModifierTooltip({
  modifier,
  ...props
}: React.ComponentProps<typeof HoverCardTrigger> & {
  modifier: Modifier
}) {
  return (
    <HoverCard openDelay={100} closeDelay={0}>
      <HoverCardTrigger {...props} />
      <HoverCardContent sideOffset={8} collisionPadding={8}>
        <div>{modifier.name}</div>
        <div className="text-muted-foreground text-xs">
          {modifier.description}
        </div>
      </HoverCardContent>
    </HoverCard>
  )
}

export { ModifierTooltip }
