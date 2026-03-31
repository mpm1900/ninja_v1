import type { Action } from '#/lib/game/action'
import { cn } from '#/lib/utils'
import type { ComponentProps } from 'react'
import { NatureBadge } from './nature-badge'

type ActionCardProps = ComponentProps<'button'> & {
  action: Action
  selected?: boolean
}

function ActionCard({
  action,
  selected = false,
  disabled = false,
  className,
  ...props
}: ActionCardProps) {
  const accuracyLabel =
    action.config.accuracy !== null ? `${action.config.accuracy}%` : '-'

  return (
    <button
      type="button"
      disabled={disabled}
      className={cn(
        'w-[200px] h-[300px] rounded-lg border-4 text-left',
        'bg-zinc-900 border-white/40 hover:border-white/60 text-zinc-100',
        'transition-all duration-200',
        'hover:-translate-y-0.5 hover:shadow-md',
        'focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-indigo-300/70',
        'flex flex-col hover:shadow-xl hover:shadow-black transition-colors',
        {
          'border-orange-400/60 hover:border-orange-400/80':
            action.config.stat === 'attack',
          'border-indigo-500/60 hover:border-indigo-400/80':
            action.config.stat === 'chakra_attack',
          'shadow-xl shadow-black': selected,
        },
        className
      )}
      {...props}
    >
      <div
        className={cn(
          'flex items-start justify-between gap-2 p-2',
          disabled && 'opacity-50'
        )}
      >
        <div className="min-w-0 flex items-center gap-1">
          {action.config.nature && (
            <NatureBadge nature={action.config.nature} />
          )}
          <div>
            <div className="truncate text-xs font-semibold text-zinc-200">
              {action.config.name}
            </div>
            <div className="text-[10px] uppercase tracking-wide text-zinc-400">
              {action.config.jutsu}
            </div>
          </div>
        </div>

        {action.config.cost && (
          <div className="text-orange-300 font-black nanum-brush-script-regular text-4xl h-8 -mt-0.5">
            {action.config.cost ?? 0}
          </div>
        )}
      </div>

      <div className="flex [&>*]:flex-1 text-[11px] border-t border-b bg-background">
        <StatChip label="Power" value={action.config.power ?? '-'} />
        <StatChip label="Acc" value={accuracyLabel} />
      </div>
    </button>
  )
}

function StatChip({ label, value }: { label: string; value: string | number }) {
  return (
    <div className="px-1.5 py-1 text-center">
      <div className="text-[9px] uppercase tracking-wide text-zinc-400">
        {label}
      </div>
      <div className="text-xs font-semibold leading-tight text-zinc-100">
        {value}
      </div>
    </div>
  )
}

export { ActionCard }
