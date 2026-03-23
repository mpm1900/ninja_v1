import { cva } from 'class-variance-authority'
import type { ClassValue } from 'class-variance-authority/types'
import { Tooltip, TooltipContent, TooltipTrigger } from './ui/tooltip'
import { natureNames, type NatureSet } from '#/lib/game/nature'

type t = Record<string, Partial<Record<NatureSet | 'none', ClassValue>>>

const variants = cva<t>(
  'text-white px-1 py-0.5 rounded text-shadow-[1px_1px_0px_#000000] border border-background',
  {
    variants: {
      variant: {
        none: '',
        pure: 'bg-slate-500',
        fire: 'bg-red-500',
        wind: 'bg-emerald-700',
        lightning: 'bg-yellow-400',
        earth: 'bg-taupe-600',
        water: 'bg-blue-500',
        yang: 'bg-neutral-300',
        yin: 'bg-violet-500',
        wood: 'bg-olive-600 shadow-[inset_0_0_0_1px_theme(colors.amber.300)] text-amber-300!',
        yinyang:
          'bg-[linear-gradient(135deg,theme(colors.neutral.300)_0%,theme(colors.neutral.300)_50%,theme(colors.violet.500)_50%,theme(colors.violet.500)_100%)] text-amber-300! shadow-[inset_0_0_0_1px_theme(colors.amber.300)]',
        jashin: 'bg-rose-950 shadow-[inset_0_0_0_1px_theme(colors.amber.300)] text-amber-300!',
      },
    },
    defaultVariants: {
      variant: 'none',
    },
  }
)

function NatureBadge({
  nature,
  className,
  ...props
}: React.ComponentProps<'span'> & { nature: NatureSet }) {
  return (
    <Tooltip>
      <TooltipTrigger asChild>
        <span
          data-role="nature"
          className={variants({ variant: nature })}
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
