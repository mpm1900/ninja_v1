import { statNames, type ActorAttackStat } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'
import type { ClassValue } from 'class-variance-authority/types'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'
import type { ComponentProps } from 'react'

type t = Record<string, Partial<Record<ActorAttackStat | 'none', ClassValue>>>
const variants = cva<t>('py-0.5 px-1 mx-px border border-lback rounded', {
  variants: {
    variant: {
      genjutsu: 'shadow-[inset_0_0_0_1px_theme(colors.rose.900)] text-rose-400',
      ninjutsu: 'shadow-[inset_0_0_0_1px_theme(colors.sky.900)] text-sky-400',
      taijutsu:
        'shadow-[inset_0_0_0_1px_theme(colors.emerald.900)] text-emerald-400',
    },
  },
})

function StatBadge({
  className,
  contentProps = {},
  stat,
  ...props
}: React.ComponentProps<'span'> & {
  stat: ActorAttackStat
  contentProps?: Partial<ComponentProps<typeof TooltipContent>>
}) {
  return (
    <Tooltip>
      <TooltipTrigger asChild>
        <span
          data-role="stat"
          className={cn(variants({ variant: stat }), className)}
          {...props}
        >
          {statNames[stat]}
        </span>
      </TooltipTrigger>
      <TooltipContent {...contentProps}>{stat}</TooltipContent>
    </Tooltip>
  )
}

export { StatBadge }
