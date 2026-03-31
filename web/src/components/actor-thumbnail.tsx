import type { Actor } from '#/lib/game/actor'
import { cn } from '#/lib/utils'
import { CircleQuestionMark } from 'lucide-react'
import type { ComponentProps } from 'react'

function ActorThumbnail({
  actor,
  hidden,
  index,
  size = 64,
  className,
  ...props
}: ComponentProps<'div'> & {
  actor: Actor
  hidden?: boolean
  index?: number
  size?: number
}) {
  const active = !!actor.position_ID
  return (
    <div
      key={actor.ID}
      style={{
        height: size,
        width: size,
      }}
      className={cn('overflow-hidden bg-card p-1 border rounded relative', {
        'bg-transparent border-transparent': index === undefined,
        'bg-foreground': active && index !== undefined,
      })}
      {...props}
    >
      <div className={cn('h-full w-full', className)}>
        {!hidden && (
          <img
            src={actor.sprite_url}
            className="h-full w-full object-cover absolute inset-0 z-10"
            width={size}
            height={size}
          />
        )}
        {!hidden && index !== undefined && (
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
        )}
        {hidden && (
          <div className="grid place-items-center h-full w-full">
            <CircleQuestionMark className="size-12 text-muted-foreground/30" />
          </div>
        )}
      </div>
    </div>
  )
}

export { ActorThumbnail }
