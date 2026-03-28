import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'

function ActorThumbnail({ actor, index }: { actor: Actor; index: number }) {
  const active = !!actor.position_ID
  return (
    <div
      key={actor.ID}
      className={cn(
        'h-18 w-18 overflow-hidden bg-card p-1 border rounded relative',
        {
          'bg-foreground': active,
        }
      )}
    >
      <img
        src={actor.sprite_url}
        className="h-full w-full object-cover absolute inset-0 z-10"
        width={64}
        height={64}
      />
      <div
        className={cn(
          'absolute -top-6 font-black text-7xl z-0 text-center text-foreground',
          {
            'text-background!': active,
          }
        )}
      >
        {index + 1}
      </div>
    </div>
  )
}

export { ActorThumbnail }
