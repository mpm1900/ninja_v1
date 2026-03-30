import { type ActorAttackStat } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { cva } from 'class-variance-authority'
import type { ClassValue } from 'class-variance-authority/types'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'
import type { ComponentProps } from 'react'

type t = Record<string, Partial<Record<ActorAttackStat | 'none', ClassValue>>>
const variants = cva<t>('py-0.5 px-1 mx-px border border-black rounded', {
  variants: {
    variant: {
      attack: 'shadow-[inset_0_0_0_1px_theme(colors.orange.900)] text-orange-400',
      chakra_attack: 'shadow-[inset_0_0_0_1px_theme(colors.indigo.900)] text-indigo-400',
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
          {stat === 'attack' && 'PH.ATK'}
          {stat === 'chakra_attack' && 'CH.ATK'}
        </span>
      </TooltipTrigger>
      <TooltipContent {...contentProps}>{stat}</TooltipContent>
    </Tooltip>
  )
}

export { StatBadge }
