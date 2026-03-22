import {
  checkActorStat,
  type Actor,
  type ActorBaseStat,
} from '#/lib/game/actor'
import { cn } from '#/lib/utils'

function getBaseStatColorClass(value: number) {
  if (value < 60) return 'text-red-500'
  if (value < 75) return 'text-orange-500'
  if (value < 90) return 'text-amber-500'
  if (value < 105) return 'text-yellow-400'
  if (value < 120) return 'text-lime-400'
  if (value < 140) return 'text-green-400'
  if (value < 160) return 'text-teal-400'
  if (value < 180) return 'text-teal-200'
  if (value < 220) return 'text-pink-200'
  return 'text-pink-300'
}

function ActorStat({
  stat,
  actor,
  showBase = true,
  ...props
}: React.ComponentProps<'span'> & {
  actor: Actor
  stat: ActorBaseStat
  showBase?: boolean
}) {
  return (
    <span data-role="actor-stat" {...props}>
      <span
        className={cn({
          'text-green-400': checkActorStat(actor, stat) === 1,
          'text-red-300': checkActorStat(actor, stat) === -1,
        })}
      >
        {actor.stats[stat]}
      </span>
      {showBase && (
        <span className="text-muted-foreground">
          {' '}
          ({actor.base_stats[stat]})
        </span>
      )}
    </span>
  )
}

function ActorStatBase({
  stat,
  actor,
  showStat = true,
  ...props
}: React.ComponentProps<'span'> & {
  actor: Actor
  stat: ActorBaseStat
  showStat?: boolean
}) {
  return (
    <span data-role="actor-stat" {...props}>
      <span className={getBaseStatColorClass(actor.base_stats[stat])}>
        {actor.base_stats[stat]}
      </span>
      {showStat && (
        <span
          className={cn('opacity-50 text-muted-foreground', {
            'text-green-400': checkActorStat(actor, stat) === 1,
            'text-red-300': checkActorStat(actor, stat) === -1,
          })}
        >
          {' '}
          ({actor.stats[stat]})
        </span>
      )}
    </span>
  )
}

export { ActorStat, ActorStatBase }
