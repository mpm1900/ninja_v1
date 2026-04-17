import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'

function ActorStatus({ actor }: { actor: Actor }) {
  return (
    <div
      className={cn(
        'absolute font-bold px-1 mx-1 left-0 right-2 h-4 text-center leading-5! rounded whitespace-nowrap -bottom-1 z-10 text-lg nanum-brush-script-regular',
        'bg-mist-300 text-background shadow-[0px_0px_5px_rgba(0,0,0,1)]',
        {
          'bg-orange-200 text-orange-900': actor.burned,
          'bg-indigo-300 text-indigo-900': actor.sleeping,
          'bg-yellow-200 text-yellow-900': actor.paralyzed,
        }
      )}
    >
      {actor.statused ? (
        <span>
          {actor.sleeping && 'SLEEP'}
          {actor.paralyzed && 'PARA'}
          {actor.burned && 'BURN'}
        </span>
      ) : (
        <span>LV {actor.level}</span>
      )}
    </div>
  )
}

export { ActorStatus }
