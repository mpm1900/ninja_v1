import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'

function ActorStatus({ actor }: { actor: Actor }) {
  return (
    <div
      className={cn(
        'absolute font-bold px-1 mx-1 left-1 right-1 h-4 text-center leading-5! rounded-xs whitespace-nowrap -bottom-1 z-10 text-lg nanum-brush-script-regular',
        'bg-stone-200 border border-stone-900 text-stone-900 shadow-[0px_1px_2px_rgba(0,0,0,1)]',
        {
          'bg-orange-200 border-orange-700 text-orange-900': actor.burned,
          'bg-indigo-300 border-indigo-700 text-indigo-900': actor.sleeping,
          'bg-yellow-200 border-yellow-700 text-yellow-900': actor.paralyzed,
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
