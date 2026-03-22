import { cva } from 'class-variance-authority'
import type { ClassValue } from 'class-variance-authority/types'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'
import { natureNames, type Nature } from '#/lib/game/actor'

type t = Record<string, Record<Nature, ClassValue>>

const variants = cva<t>('text-white px-1 py-0.5 rounded text-shadow-md', {
  variants: {
    variant: {
      fire: 'bg-red-500',
      wind: 'bg-emerald-700',
      lightning: 'bg-yellow-400 text-black!',
      earth: 'bg-taupe-600',
      water: 'bg-blue-500',
      yang: 'bg-neutral-300 text-black!',
      yin: 'bg-violet-500',
    },
  },
  defaultVariants: {
    variant: 'fire',
  },
})

function NatureBadge({
  nature,
  className,
  ...props
}: React.ComponentProps<'span'> & { nature: Nature }) {
  return (
    <Tooltip>
      <TooltipTrigger asChild>
        <span
          data-role="nature"
          className={variants({ variant: nature }) + ''}
          {...props}
        >
          {natureNames[nature] ?? nature}
        </span>
      </TooltipTrigger>
      <TooltipContent>{nature}</TooltipContent>
    </Tooltip>
  )
}

export { NatureBadge }
