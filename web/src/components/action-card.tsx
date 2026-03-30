import type { Action } from '#/lib/game/action'
import { statNames } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import type { ComponentProps } from 'react'
import { NatureBadge } from './nature-badge'

function StatBar({ action }: { action: Action }) {
  const stat = action.config.stat!
  return (
    <div
      className={cn('', {
        'text-sky-400': stat === 'chakra_attack',
        'text-emerald-400': stat === 'attack',
      })}
    >
      <div>
        {statNames[stat]}:{' '}
        {action.config.power ? `POW: ${action.config.power}` : '-'}{' '}
        {action.config.accuracy ? `ACC: ${action.config.accuracy}%` : '-'}
      </div>
    </div>
  )
}

function ActionCard({
  action,
  className,
  ...props
}: ComponentProps<'div'> & { action: Action }) {
  return (
    <div
      className={cn(
        'p-2 border-4 border-neutral-600 bg-input/40 rounded-lg w-[240px] h-[360px]',
        {
          'border-green-700': action.config.stat === 'attack',
          'border-cyan-600': action.config.stat === 'chakra_attack'
        },
        className
      )}
      {...props}
    >
      <div className="flex gap-4 justify-between">
        <div className="flex gap-1">
          {action.config.nature && (
            <div>
              <NatureBadge nature={action.config.nature} />
            </div>
          )}
          <div>{action.config.name}</div>
        </div>
        <div>
          <div className="bg-orange-400/70 size-7 font-black grid place-items-center text-xs rounded-full border border-black text-shadow-[1px_1px_0px_#000000] shadow-[inset_0_0_0_1px_theme(colors.neutral.400)] ">
            {action.config.cost ?? 0}
          </div>
        </div>
      </div>
      <div className="h-20"></div>
      {action.config.stat && <StatBar action={action} />}
    </div>
  )
}

export { ActionCard }
