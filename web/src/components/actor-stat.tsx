import {
  checkActorStat,
  type Actor,
  type ActorBaseStat,
} from '#/lib/game/actor'
import { cn } from '#/lib/utils'

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

export { ActorStat }
